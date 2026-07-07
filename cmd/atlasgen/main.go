package main

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"flag"
	"fmt"
	"html"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	atlasPath              = "data/atlas/ATLAS.yaml"
	resourcesContentDir    = "content/resources"
	defaultTranslationPath = "data/translations/ru.yaml"
	reportPath             = "reports/translation-coverage.md"
	reportPagePath         = "content/translation-coverage.md"
	reportDataPath         = "data/generated/translation-report.yaml"
	generatedAtlasDataPath = "data/generated/atlas.yaml"
	generatedBy            = "atlasgen"
)

type Atlas struct {
	ID          string   `yaml:"id"`
	Name        string   `yaml:"name"`
	Version     string   `yaml:"version"`
	Matrices    []Matrix `yaml:"matrices"`
	CaseStudies []Study  `yaml:"case-studies"`
}

type Matrix struct {
	ID          string       `yaml:"id"`
	Name        string       `yaml:"name"`
	Tactics     []Object     `yaml:"tactics"`
	Techniques  []Object     `yaml:"techniques"`
	Mitigations []Mitigation `yaml:"mitigations"`
}

type Object struct {
	ID              string           `yaml:"id"`
	Name            string           `yaml:"name"`
	Description     string           `yaml:"description"`
	ObjectType      string           `yaml:"object-type"`
	AttackReference *AttackReference `yaml:"ATT&CK-reference"`
	Tactics         []string         `yaml:"tactics"`
	SubtechniqueOf  string           `yaml:"subtechnique-of"`
	References      []Reference      `yaml:"references"`
	CreatedDate     string           `yaml:"created_date"`
	ModifiedDate    string           `yaml:"modified_date"`
	Maturity        string           `yaml:"maturity"`
	Platforms       []string         `yaml:"platforms,omitempty"`
}

type Mitigation struct {
	ID              string           `yaml:"id"`
	Name            string           `yaml:"name"`
	Description     string           `yaml:"description"`
	ObjectType      string           `yaml:"object-type"`
	AttackReference *AttackReference `yaml:"ATT&CK-reference"`
	Techniques      []TechniqueUse   `yaml:"techniques"`
	MLLifecycle     StringList       `yaml:"ml-lifecycle"`
	Category        []string         `yaml:"category"`
	CreatedDate     string           `yaml:"created_date"`
	ModifiedDate    string           `yaml:"modified_date"`
}

func (m *Mitigation) UnmarshalYAML(value *yaml.Node) error {
	type rawMitigation struct {
		ID              string           `yaml:"id"`
		Name            string           `yaml:"name"`
		Description     string           `yaml:"description"`
		ObjectType      string           `yaml:"object-type"`
		AttackReference *AttackReference `yaml:"ATT&CK-reference"`
		Techniques      []TechniqueUse   `yaml:"techniques"`
		MLLifecycle     yaml.Node        `yaml:"ml-lifecycle"`
		Category        []string         `yaml:"category"`
		CreatedDate     string           `yaml:"created_date"`
		ModifiedDate    string           `yaml:"modified_date"`
	}
	var raw rawMitigation
	if err := value.Decode(&raw); err != nil {
		return err
	}
	lifecycle, extraTechniques := parseMLLifecycle(raw.MLLifecycle)
	*m = Mitigation{
		ID:              raw.ID,
		Name:            raw.Name,
		Description:     raw.Description,
		ObjectType:      raw.ObjectType,
		AttackReference: raw.AttackReference,
		Techniques:      append(raw.Techniques, extraTechniques...),
		MLLifecycle:     lifecycle,
		Category:        raw.Category,
		CreatedDate:     raw.CreatedDate,
		ModifiedDate:    raw.ModifiedDate,
	}
	return nil
}

func parseMLLifecycle(value yaml.Node) (StringList, []TechniqueUse) {
	if value.Kind == 0 {
		return nil, nil
	}
	if value.Kind != yaml.SequenceNode {
		text := strings.TrimSpace(value.Value)
		if text == "" {
			return nil, nil
		}
		return StringList{text}, nil
	}
	lifecycle := make([]string, 0, len(value.Content))
	techniques := []TechniqueUse{}
	for _, item := range value.Content {
		switch item.Kind {
		case yaml.ScalarNode:
			text := strings.TrimSpace(item.Value)
			if text != "" {
				lifecycle = append(lifecycle, text)
			}
		case yaml.MappingNode:
			var technique TechniqueUse
			if err := item.Decode(&technique); err == nil && strings.TrimSpace(technique.ID) != "" {
				techniques = append(techniques, technique)
			}
		}
	}
	return lifecycle, techniques
}

type StringList []string

func (l *StringList) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.SequenceNode {
		var single string
		if err := value.Decode(&single); err != nil {
			return err
		}
		if strings.TrimSpace(single) != "" {
			*l = []string{single}
		}
		return nil
	}
	items := make([]string, 0, len(value.Content))
	for _, item := range value.Content {
		if item.Kind != yaml.ScalarNode {
			continue
		}
		text := strings.TrimSpace(item.Value)
		if text != "" {
			items = append(items, text)
		}
	}
	*l = items
	return nil
}

type TechniqueUse struct {
	ID  string `yaml:"id"`
	Use string `yaml:"use"`
}

type Study struct {
	ID                      string      `yaml:"id"`
	Name                    string      `yaml:"name"`
	ObjectType              string      `yaml:"object-type"`
	Summary                 string      `yaml:"summary"`
	IncidentDate            string      `yaml:"incident-date"`
	IncidentDateGranularity string      `yaml:"incident-date-granularity"`
	Procedure               []Procedure `yaml:"procedure"`
	Target                  string      `yaml:"target"`
	Actor                   string      `yaml:"actor"`
	Reporter                string      `yaml:"reporter"`
	CaseStudyType           string      `yaml:"case-study-type"`
	References              []Reference `yaml:"references"`
}

type Procedure struct {
	Tactic      string `yaml:"tactic"`
	Technique   string `yaml:"technique"`
	Description string `yaml:"description"`
}

type Reference struct {
	Title string `yaml:"title"`
	URL   string `yaml:"url"`
}

type AttackReference struct {
	ID  string `yaml:"id"`
	URL string `yaml:"url"`
}

type AtlasV6 struct {
	FormatVersion string                       `yaml:"format-version"`
	Collection    CollectionV6                 `yaml:"collection"`
	Matrix        MatrixV6                     `yaml:"matrix"`
	Tactics       map[string]ObjectV6          `yaml:"tactics"`
	Techniques    map[string]ObjectV6          `yaml:"techniques"`
	Mitigations   map[string]MitigationV6      `yaml:"mitigations"`
	CaseStudies   map[string]StudyV6           `yaml:"case-studies"`
	Relationships map[string]RelationshipSetV6 `yaml:"relationships"`
}

type CollectionV6 struct {
	ID          string `yaml:"id"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Version     string `yaml:"version"`
}

type MatrixV6 struct {
	ID          string `yaml:"id"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type ObjectV6 struct {
	ID              string           `yaml:"id"`
	Name            string           `yaml:"name"`
	Description     string           `yaml:"description"`
	ObjectType      string           `yaml:"object-type"`
	AttackReference *AttackReference `yaml:"attack-reference"`
	References      []Reference      `yaml:"references"`
	CreatedDate     string           `yaml:"created-date"`
	ModifiedDate    string           `yaml:"modified-date"`
	Maturity        string           `yaml:"maturity"`
	Platforms       []string         `yaml:"platforms"`
}

type MitigationV6 struct {
	ID              string           `yaml:"id"`
	Name            string           `yaml:"name"`
	Description     string           `yaml:"description"`
	ObjectType      string           `yaml:"object-type"`
	AttackReference *AttackReference `yaml:"attack-reference"`
	References      []Reference      `yaml:"references"`
	CreatedDate     string           `yaml:"created-date"`
	ModifiedDate    string           `yaml:"modified-date"`
	LifecyclePhases StringList       `yaml:"lifecycle-phases"`
	Categories      []string         `yaml:"categories"`
}

type StudyV6 struct {
	ID                      string      `yaml:"id"`
	Name                    string      `yaml:"name"`
	ObjectType              string      `yaml:"object-type"`
	Description             string      `yaml:"description"`
	Summary                 string      `yaml:"summary"`
	IncidentDate            string      `yaml:"incident-date"`
	IncidentDateGranularity string      `yaml:"incident-date-granularity"`
	Procedure               []Procedure `yaml:"procedure"`
	Target                  string      `yaml:"target"`
	Actor                   string      `yaml:"actor"`
	Reporter                string      `yaml:"reporter"`
	CaseStudyType           string      `yaml:"case-study-type"`
	References              []Reference `yaml:"references"`
}

type RelationshipSetV6 struct {
	Achieves    []RelationshipV6 `yaml:"achieves"`
	Specializes []RelationshipV6 `yaml:"specializes"`
	Mitigates   []RelationshipV6 `yaml:"mitigates"`
	Employs     []RelationshipV6 `yaml:"employs"`
	Sequences   []RelationshipV6 `yaml:"sequences"`
}

type RelationshipV6 struct {
	Source           string   `yaml:"source"`
	Target           string   `yaml:"target"`
	RelationshipType string   `yaml:"relationship-type"`
	Description      string   `yaml:"description"`
	Tactic           string   `yaml:"tactic"`
	StepID           string   `yaml:"step-id"`
	Position         int      `yaml:"position"`
	LeadsTo          []string `yaml:"leads-to"`
}

type Translations struct {
	TechniqueUses []TechniqueUseTextTranslation `yaml:"technique_uses"`
	Objects       map[string]Translation        `yaml:"objects"`
}

type Translation struct {
	Name        string                 `yaml:"name"`
	Description string                 `yaml:"description"`
	Summary     string                 `yaml:"summary"`
	Procedure   []ProcedureTranslation `yaml:"procedure"`
	Techniques  []TechniqueTranslation `yaml:"techniques"`
	Source      Source                 `yaml:"source"`
}

type TechniqueTranslation struct {
	ID     string `yaml:"id"`
	Use    string `yaml:"use"`
	Source Source `yaml:"source"`
}

type TechniqueUseTextTranslation struct {
	Source string `yaml:"source"`
	Use    string `yaml:"use"`
}

type ProcedureTranslation struct {
	Tactic      string `yaml:"tactic"`
	Technique   string `yaml:"technique"`
	Description string `yaml:"description"`
	Source      Source `yaml:"source"`
}

type Source struct {
	NameSHA256        string `yaml:"name_sha256"`
	DescriptionSHA256 string `yaml:"description_sha256"`
	SummarySHA256     string `yaml:"summary_sha256"`
	UseSHA256         string `yaml:"use_sha256"`
}

type Page struct {
	Section     string
	Kind        string
	ID          string
	SourceName  string
	Title       string
	Description string
	Body        string
	Params      map[string]any
}

type CoverageItem struct {
	ID                      string   `yaml:"id"`
	Name                    string   `yaml:"name"`
	Type                    string   `yaml:"type"`
	HasName                 bool     `yaml:"has_name"`
	HasBody                 bool     `yaml:"has_body"`
	SourceField             string   `yaml:"source_field"`
	TotalProcedures         int      `yaml:"total_procedures,omitempty"`
	TranslatedProcedures    int      `yaml:"translated_procedures,omitempty"`
	ProcedureNeedsReview    int      `yaml:"procedure_needs_review,omitempty"`
	NeedsReview             bool     `yaml:"needs_review,omitempty"`
	ReviewFields            []string `yaml:"review_fields,omitempty"`
	SourceSHA256            string   `yaml:"source_sha256,omitempty"`
	TranslationSourceSHA256 string   `yaml:"translation_source_sha256,omitempty"`
}

type CoverageSummary struct {
	Version string                 `yaml:"version"`
	Totals  map[string]TypeSummary `yaml:"totals"`
	Missing []CoverageItem         `yaml:"missing"`
	Partial []CoverageItem         `yaml:"partial"`
	Review  []CoverageItem         `yaml:"review"`
}

type TypeSummary struct {
	Total                int `yaml:"total"`
	NameTranslated       int `yaml:"name_translated"`
	BodyTranslated       int `yaml:"body_translated"`
	ProceduresTotal      int `yaml:"procedures_total,omitempty"`
	ProceduresTranslated int `yaml:"procedures_translated,omitempty"`
	FullyTranslated      int `yaml:"fully_translated"`
	NeedsReview          int `yaml:"needs_review"`
}

type Catalog struct {
	Translations           Translations
	Tactics                map[string]Object
	Techniques             map[string]Object
	Mitigations            map[string]Mitigation
	Studies                map[string]Study
	TacticTechniques       map[string][]Object
	SubtechniquesByParent  map[string][]Object
	MitigationsByTechnique map[string][]MitigationUseEntry
	ProceduresByTechnique  map[string][]ProcedureExample
	ProceduresByTactic     map[string][]ProcedureExample
}

type MitigationUseEntry struct {
	Mitigation  Mitigation
	TechniqueID string
	Use         string
}

type ProcedureExample struct {
	Study     Study
	Procedure Procedure
	Index     int
}

func main() {
	if err := runCommand(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "atlasgen: %v\n", err)
		os.Exit(1)
	}
}

func runCommand(args []string) error {
	flags := flag.NewFlagSet("atlasgen", flag.ContinueOnError)
	sourceHashID := flags.String("source-hash", "", "print source hash YAML snippet for an ATLAS object ID")
	procedureHashesID := flags.String("procedure-hashes", "", "print source hash YAML snippets for a case study's procedures")
	techniqueUseHashesID := flags.String("technique-use-hashes", "", "print source hash YAML snippets for a mitigation's related technique uses")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if strings.TrimSpace(*techniqueUseHashesID) != "" {
		return printTechniqueUseHashes(*techniqueUseHashesID)
	}
	if strings.TrimSpace(*procedureHashesID) != "" {
		return printProcedureHashes(*procedureHashesID)
	}
	if strings.TrimSpace(*sourceHashID) != "" {
		return printSourceHash(*sourceHashID)
	}
	return run()
}

func run() error {
	atlas, err := readAtlas(atlasPath)
	if err != nil {
		return err
	}
	translations, err := readYAML[Translations](translationFilePath())
	if err != nil {
		return err
	}
	if len(atlas.Matrices) == 0 {
		return errors.New("ATLAS.yaml has no matrices")
	}
	if translations.Objects == nil {
		translations.Objects = map[string]Translation{}
	}

	matrix := atlas.Matrices[0]
	catalog := buildCatalog(matrix, atlas.CaseStudies, translations)
	pages := make([]Page, 0, len(matrix.Tactics)+len(matrix.Techniques)+len(matrix.Mitigations)+len(atlas.CaseStudies))
	coverage := newCoverage(atlas.Version)

	for _, obj := range matrix.Tactics {
		pages = append(pages, pageForTactic(obj, catalog))
		coverage.add("tactic", obj.ID, obj.Name, "description", obj.Description, translations.Objects[obj.ID])
	}
	for _, obj := range matrix.Techniques {
		pages = append(pages, pageForTechnique(obj, catalog))
		coverage.add("technique", obj.ID, obj.Name, "description", obj.Description, translations.Objects[obj.ID])
	}
	for _, obj := range matrix.Mitigations {
		pages = append(pages, pageForMitigation(obj, catalog))
		coverage.add("mitigation", obj.ID, obj.Name, "description", obj.Description, translations.Objects[obj.ID])
	}
	for _, obj := range atlas.CaseStudies {
		pages = append(pages, pageForStudy(obj, catalog))
		coverage.addCaseStudy(obj, translations.Objects[obj.ID])
	}
	if err := coverage.addResources(resourcesContentDir); err != nil {
		return err
	}

	if err := writePages(pages); err != nil {
		return err
	}
	if err := writeCoverageReports(coverage); err != nil {
		return err
	}
	if err := writeGeneratedAtlasData(atlas); err != nil {
		return err
	}

	fmt.Printf("Generated %d pages from ATLAS %s.\n", len(pages), atlas.Version)
	fmt.Printf("Translation coverage report: %s\n", reportPath)
	return nil
}

func translationFilePath() string {
	if path := strings.TrimSpace(os.Getenv("ATLAS_TRANSLATION_FILE")); path != "" {
		return path
	}
	if lang := strings.TrimSpace(os.Getenv("ATLAS_LANG")); lang != "" {
		return filepath.Join("data", "translations", lang+".yaml")
	}
	return defaultTranslationPath
}

func printSourceHash(id string) error {
	atlas, err := readAtlas(atlasPath)
	if err != nil {
		return err
	}
	nameHash, field, hash, err := sourceHashesForID(atlas, id)
	if err != nil {
		return err
	}
	fmt.Printf("source:\n  name_sha256: %q\n  %s_sha256: %q\n", nameHash, field, hash)
	return nil
}

func printProcedureHashes(id string) error {
	atlas, err := readAtlas(atlasPath)
	if err != nil {
		return err
	}
	id = strings.TrimSpace(id)
	for _, study := range atlas.CaseStudies {
		if study.ID != id {
			continue
		}
		fmt.Println("procedure:")
		for _, step := range study.Procedure {
			fmt.Printf("  - tactic: %s\n", step.Tactic)
			fmt.Printf("    technique: %s\n", step.Technique)
			fmt.Printf("    source:\n      description_sha256: %q\n", sourceHash(step.Description))
		}
		return nil
	}
	return fmt.Errorf("case study %q not found", id)
}

func printTechniqueUseHashes(id string) error {
	atlas, err := readAtlas(atlasPath)
	if err != nil {
		return err
	}
	id = strings.TrimSpace(id)
	for _, matrix := range atlas.Matrices {
		for _, mitigation := range matrix.Mitigations {
			if mitigation.ID != id {
				continue
			}
			fmt.Println("techniques:")
			for _, technique := range mitigation.Techniques {
				fmt.Printf("  - id: %s\n", technique.ID)
				fmt.Printf("    source:\n      use_sha256: %q\n", sourceHash(technique.Use))
			}
			return nil
		}
	}
	return fmt.Errorf("mitigation %q not found", id)
}

func sourceHashForID(atlas Atlas, id string) (string, string, error) {
	_, field, hash, err := sourceHashesForID(atlas, id)
	return field, hash, err
}

func sourceHashesForID(atlas Atlas, id string) (string, string, string, error) {
	id = strings.TrimSpace(id)
	for _, matrix := range atlas.Matrices {
		for _, obj := range matrix.Tactics {
			if obj.ID == id {
				return sourceHash(obj.Name), "description", sourceHash(obj.Description), nil
			}
		}
		for _, obj := range matrix.Techniques {
			if obj.ID == id {
				return sourceHash(obj.Name), "description", sourceHash(obj.Description), nil
			}
		}
		for _, obj := range matrix.Mitigations {
			if obj.ID == id {
				return sourceHash(obj.Name), "description", sourceHash(obj.Description), nil
			}
		}
	}
	for _, study := range atlas.CaseStudies {
		if study.ID == id {
			return sourceHash(study.Name), "summary", sourceHash(study.Summary), nil
		}
	}
	return "", "", "", fmt.Errorf("source object %q not found", id)
}

func readYAML[T any](path string) (T, error) {
	var value T
	data, err := os.ReadFile(path)
	if err != nil {
		return value, err
	}
	if err := yaml.Unmarshal(data, &value); err != nil {
		return value, fmt.Errorf("read %s: %w", path, err)
	}
	return value, nil
}

func readAtlas(path string) (Atlas, error) {
	var atlas Atlas
	data, err := os.ReadFile(path)
	if err != nil {
		return atlas, err
	}
	var probe struct {
		FormatVersion string `yaml:"format-version"`
	}
	if err := yaml.Unmarshal(data, &probe); err != nil {
		return atlas, fmt.Errorf("read %s: %w", path, err)
	}
	if strings.TrimSpace(probe.FormatVersion) == "" {
		if err := yaml.Unmarshal(data, &atlas); err != nil {
			return atlas, fmt.Errorf("read %s: %w", path, err)
		}
		return atlas, nil
	}
	var atlasV6 AtlasV6
	if err := yaml.Unmarshal(data, &atlasV6); err != nil {
		return atlas, fmt.Errorf("read %s: %w", path, err)
	}
	return adaptAtlasV6(atlasV6), nil
}

func adaptAtlasV6(source AtlasV6) Atlas {
	matrix := Matrix{
		ID:   firstNonEmpty(source.Matrix.ID, "ATLAS"),
		Name: firstNonEmpty(source.Matrix.Name, source.Collection.Name),
	}
	atlas := Atlas{
		ID:      firstNonEmpty(source.Collection.ID, "ATLAS"),
		Name:    firstNonEmpty(source.Collection.Description, source.Collection.Name),
		Version: firstNonEmpty(source.Collection.Version, source.FormatVersion),
	}

	tacticIDs := orderedRelationshipTargets(source.Relationships[matrix.ID].Sequences)
	if len(tacticIDs) == 0 {
		tacticIDs = orderedRelationshipTargets(source.Relationships["ATLAS-matrix"].Sequences)
	}
	if len(tacticIDs) == 0 {
		tacticIDs = sortedKeys(source.Tactics)
	}
	for _, id := range tacticIDs {
		if tactic, ok := source.Tactics[id]; ok {
			matrix.Tactics = append(matrix.Tactics, objectFromV6(tactic))
		}
	}

	for _, id := range sortedKeys(source.Techniques) {
		technique := objectFromV6(source.Techniques[id])
		for _, rel := range source.Relationships[id].Achieves {
			if strings.TrimSpace(rel.Target) != "" {
				technique.Tactics = append(technique.Tactics, strings.TrimSpace(rel.Target))
			}
		}
		for _, rel := range source.Relationships[id].Specializes {
			if strings.TrimSpace(rel.Target) != "" {
				technique.SubtechniqueOf = strings.TrimSpace(rel.Target)
				break
			}
		}
		matrix.Techniques = append(matrix.Techniques, technique)
	}

	for _, id := range sortedKeys(source.Mitigations) {
		mitigation := mitigationFromV6(source.Mitigations[id])
		for _, rel := range source.Relationships[id].Mitigates {
			if strings.TrimSpace(rel.Target) == "" {
				continue
			}
			mitigation.Techniques = append(mitigation.Techniques, TechniqueUse{
				ID:  strings.TrimSpace(rel.Target),
				Use: strings.TrimSpace(rel.Description),
			})
		}
		matrix.Mitigations = append(matrix.Mitigations, mitigation)
	}

	for _, id := range sortedKeys(source.CaseStudies) {
		study := studyFromV6(source.CaseStudies[id])
		if len(study.Procedure) == 0 {
			relationships := append([]RelationshipV6{}, source.Relationships[id].Employs...)
			sort.SliceStable(relationships, func(i, j int) bool {
				return relationshipOrder(relationships[i]) < relationshipOrder(relationships[j])
			})
			for _, rel := range relationships {
				if strings.TrimSpace(rel.Target) == "" {
					continue
				}
				study.Procedure = append(study.Procedure, Procedure{
					Tactic:      strings.TrimSpace(rel.Tactic),
					Technique:   strings.TrimSpace(rel.Target),
					Description: strings.TrimSpace(rel.Description),
				})
			}
		}
		atlas.CaseStudies = append(atlas.CaseStudies, study)
	}

	atlas.Matrices = []Matrix{matrix}
	return atlas
}

func objectFromV6(source ObjectV6) Object {
	return Object{
		ID:              strings.TrimSpace(source.ID),
		Name:            strings.TrimSpace(source.Name),
		Description:     strings.TrimSpace(source.Description),
		ObjectType:      strings.TrimSpace(source.ObjectType),
		AttackReference: source.AttackReference,
		References:      source.References,
		CreatedDate:     strings.TrimSpace(source.CreatedDate),
		ModifiedDate:    strings.TrimSpace(source.ModifiedDate),
		Maturity:        normalizeMaturity(source.Maturity),
		Platforms:       cleanStrings(source.Platforms),
	}
}

func mitigationFromV6(source MitigationV6) Mitigation {
	return Mitigation{
		ID:              strings.TrimSpace(source.ID),
		Name:            strings.TrimSpace(source.Name),
		Description:     strings.TrimSpace(source.Description),
		ObjectType:      strings.TrimSpace(source.ObjectType),
		AttackReference: source.AttackReference,
		MLLifecycle:     source.LifecyclePhases,
		Category:        cleanStrings(source.Categories),
		CreatedDate:     strings.TrimSpace(source.CreatedDate),
		ModifiedDate:    strings.TrimSpace(source.ModifiedDate),
	}
}

func studyFromV6(source StudyV6) Study {
	return Study{
		ID:                      strings.TrimSpace(source.ID),
		Name:                    strings.TrimSpace(source.Name),
		ObjectType:              strings.TrimSpace(source.ObjectType),
		Summary:                 firstNonEmpty(source.Summary, source.Description),
		IncidentDate:            strings.TrimSpace(source.IncidentDate),
		IncidentDateGranularity: strings.TrimSpace(source.IncidentDateGranularity),
		Procedure:               source.Procedure,
		Target:                  strings.TrimSpace(source.Target),
		Actor:                   strings.TrimSpace(source.Actor),
		Reporter:                strings.TrimSpace(source.Reporter),
		CaseStudyType:           strings.TrimSpace(source.CaseStudyType),
		References:              source.References,
	}
}

func orderedRelationshipTargets(relationships []RelationshipV6) []string {
	ordered := append([]RelationshipV6{}, relationships...)
	sort.SliceStable(ordered, func(i, j int) bool {
		return relationshipOrder(ordered[i]) < relationshipOrder(ordered[j])
	})
	targets := make([]string, 0, len(ordered))
	seen := map[string]bool{}
	for _, rel := range ordered {
		target := strings.TrimSpace(rel.Target)
		if target == "" || seen[target] {
			continue
		}
		seen[target] = true
		targets = append(targets, target)
	}
	return targets
}

func relationshipOrder(rel RelationshipV6) string {
	if rel.Position > 0 {
		return fmt.Sprintf("%08d", rel.Position)
	}
	if strings.TrimSpace(rel.StepID) != "" {
		return strings.TrimSpace(rel.StepID)
	}
	return strings.TrimSpace(rel.Target)
}

func sortedKeys[T any](values map[string]T) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if text := strings.TrimSpace(value); text != "" {
			return text
		}
	}
	return ""
}

func cleanStrings(values []string) []string {
	cleaned := make([]string, 0, len(values))
	seen := map[string]bool{}
	for _, value := range values {
		text := strings.TrimSpace(value)
		if text == "" || seen[text] {
			continue
		}
		seen[text] = true
		cleaned = append(cleaned, text)
	}
	return cleaned
}

func normalizeMaturity(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "feasible":
		return "feasible"
	case "demonstrated":
		return "demonstrated"
	case "realized":
		return "realized"
	default:
		return strings.TrimSpace(value)
	}
}

func buildCatalog(matrix Matrix, studies []Study, translations Translations) Catalog {
	c := Catalog{
		Translations:           translations,
		Tactics:                map[string]Object{},
		Techniques:             map[string]Object{},
		Mitigations:            map[string]Mitigation{},
		Studies:                map[string]Study{},
		TacticTechniques:       map[string][]Object{},
		SubtechniquesByParent:  map[string][]Object{},
		MitigationsByTechnique: map[string][]MitigationUseEntry{},
		ProceduresByTechnique:  map[string][]ProcedureExample{},
		ProceduresByTactic:     map[string][]ProcedureExample{},
	}

	for _, tactic := range matrix.Tactics {
		c.Tactics[tactic.ID] = tactic
	}
	for _, technique := range matrix.Techniques {
		c.Techniques[technique.ID] = technique
		if technique.SubtechniqueOf != "" {
			c.SubtechniquesByParent[technique.SubtechniqueOf] = append(c.SubtechniquesByParent[technique.SubtechniqueOf], technique)
		}
	}
	for _, technique := range matrix.Techniques {
		if len(technique.Tactics) == 0 {
			continue
		}
		for _, tacticID := range technique.Tactics {
			c.TacticTechniques[tacticID] = append(c.TacticTechniques[tacticID], technique)
			for _, subtechnique := range c.SubtechniquesByParent[technique.ID] {
				c.TacticTechniques[tacticID] = append(c.TacticTechniques[tacticID], subtechnique)
			}
		}
	}
	for _, mitigation := range matrix.Mitigations {
		c.Mitigations[mitigation.ID] = mitigation
		for _, technique := range mitigation.Techniques {
			entry := MitigationUseEntry{Mitigation: mitigation, TechniqueID: technique.ID, Use: technique.Use}
			c.MitigationsByTechnique[technique.ID] = append(c.MitigationsByTechnique[technique.ID], entry)
		}
	}
	for _, study := range studies {
		c.Studies[study.ID] = study
		for i, procedure := range study.Procedure {
			example := ProcedureExample{Study: study, Procedure: procedure, Index: i}
			c.ProceduresByTechnique[procedure.Technique] = append(c.ProceduresByTechnique[procedure.Technique], example)
			c.ProceduresByTactic[procedure.Tactic] = append(c.ProceduresByTactic[procedure.Tactic], example)
		}
	}

	for key := range c.TacticTechniques {
		sortObjects(c.TacticTechniques[key])
	}
	for key := range c.SubtechniquesByParent {
		sortObjects(c.SubtechniquesByParent[key])
	}
	for key := range c.MitigationsByTechnique {
		sort.Slice(c.MitigationsByTechnique[key], func(i, j int) bool {
			return c.MitigationsByTechnique[key][i].Mitigation.ID < c.MitigationsByTechnique[key][j].Mitigation.ID
		})
	}
	for key := range c.ProceduresByTechnique {
		sortProcedureExamples(c.ProceduresByTechnique[key])
	}
	for key := range c.ProceduresByTactic {
		sortProcedureExamples(c.ProceduresByTactic[key])
	}
	return c
}

func pageForTactic(obj Object, catalog Catalog) Page {
	tr := catalog.Translations.Objects[obj.ID]
	body := translatedBody(obj.Description, tr.Description)
	var b strings.Builder
	writeStatus(&b, tr.Description)
	writeParagraphs(&b, body)
	if techniques := catalog.TacticTechniques[obj.ID]; len(techniques) > 0 {
		writeObjectRelations(&b, "Техники", "techniques", techniques, catalog)
	}
	if examples := catalog.ProceduresByTactic[obj.ID]; len(examples) > 0 {
		writeProcedureExamples(&b, examples, catalog, 12)
	}
	return Page{
		Section:     "tactics",
		Kind:        "tactic",
		ID:          obj.ID,
		SourceName:  obj.Name,
		Title:       catalog.objectName(obj.ID, obj.Name),
		Description: body,
		Body:        b.String(),
		Params: map[string]any{
			"created_date":    obj.CreatedDate,
			"modified_date":   obj.ModifiedDate,
			"attack_ref_id":   attackRefID(obj.AttackReference),
			"attack_ref_url":  attackRefURL(obj.AttackReference),
			"technique_count": len(catalog.TacticTechniques[obj.ID]),
			"procedure_count": len(catalog.ProceduresByTactic[obj.ID]),
		},
	}
}

func pageForTechnique(obj Object, catalog Catalog) Page {
	tr := catalog.Translations.Objects[obj.ID]
	body := translatedBody(obj.Description, tr.Description)
	var b strings.Builder
	writeStatus(&b, tr.Description)
	writeParagraphs(&b, body)
	if len(obj.Tactics) > 0 {
		writeTacticRelations(&b, "Тактики", obj.Tactics, catalog)
	}
	if obj.SubtechniqueOf != "" {
		parent := catalog.Techniques[obj.SubtechniqueOf]
		writeObjectRelations(&b, "Родительская техника", "techniques", []Object{parent}, catalog)
	}
	if subtechniques := catalog.SubtechniquesByParent[obj.ID]; len(subtechniques) > 0 {
		writeObjectRelations(&b, "Подтехники", "techniques", subtechniques, catalog)
	}
	if obj.SubtechniqueOf != "" {
		if siblings := catalog.SubtechniquesByParent[obj.SubtechniqueOf]; len(siblings) > 1 {
			filtered := make([]Object, 0, len(siblings)-1)
			for _, sibling := range siblings {
				if sibling.ID == obj.ID {
					continue
				}
				filtered = append(filtered, sibling)
			}
			writeObjectRelations(&b, "Другие подтехники родителя", "techniques", filtered, catalog)
		}
	}
	if mitigations := catalog.MitigationsByTechnique[obj.ID]; len(mitigations) > 0 {
		writeMitigationRelations(&b, mitigations, catalog)
	}
	if examples := catalog.ProceduresByTechnique[obj.ID]; len(examples) > 0 {
		writeProcedureExamples(&b, examples, catalog, 0)
	}
	writeReferences(&b, obj.References)
	return Page{
		Section:     "techniques",
		Kind:        "technique",
		ID:          obj.ID,
		SourceName:  obj.Name,
		Title:       catalog.objectName(obj.ID, obj.Name),
		Description: body,
		Body:        b.String(),
		Params: map[string]any{
			"created_date":       obj.CreatedDate,
			"modified_date":      obj.ModifiedDate,
			"maturity":           obj.Maturity,
			"platforms":          obj.Platforms,
			"tactics":            obj.Tactics,
			"subtechnique_of":    obj.SubtechniqueOf,
			"attack_ref_id":      attackRefID(obj.AttackReference),
			"attack_ref_url":     attackRefURL(obj.AttackReference),
			"subtechnique_count": len(catalog.SubtechniquesByParent[obj.ID]),
			"mitigation_count":   len(catalog.MitigationsByTechnique[obj.ID]),
			"procedure_count":    len(catalog.ProceduresByTechnique[obj.ID]),
		},
	}
}

func pageForMitigation(obj Mitigation, catalog Catalog) Page {
	tr := catalog.Translations.Objects[obj.ID]
	body := translatedBody(obj.Description, tr.Description)
	var b strings.Builder
	writeStatus(&b, tr.Description)
	writeParagraphs(&b, body)
	if len(obj.Techniques) > 0 {
		writeTechniqueUseRelations(&b, obj.ID, obj.Techniques, catalog)
	}
	return Page{
		Section:     "mitigations",
		Kind:        "mitigation",
		ID:          obj.ID,
		SourceName:  obj.Name,
		Title:       catalog.objectName(obj.ID, obj.Name),
		Description: body,
		Body:        b.String(),
		Params: map[string]any{
			"created_date":    obj.CreatedDate,
			"modified_date":   obj.ModifiedDate,
			"category":        obj.Category,
			"ml_lifecycle":    obj.MLLifecycle,
			"attack_ref_id":   attackRefID(obj.AttackReference),
			"attack_ref_url":  attackRefURL(obj.AttackReference),
			"technique_count": len(obj.Techniques),
		},
	}
}

func pageForStudy(obj Study, catalog Catalog) Page {
	tr := catalog.Translations.Objects[obj.ID]
	body := translatedBody(obj.Summary, tr.Summary)
	hasSanitizedPayloadExamples := hasRawHTMLPayloadBlock(body)
	var b strings.Builder
	writeStatus(&b, tr.Summary)
	writeParagraphs(&b, body)
	procedure := make([]map[string]string, 0, len(obj.Procedure))
	for i, step := range obj.Procedure {
		tactic := catalog.Tactics[step.Tactic]
		technique := catalog.Techniques[step.Technique]
		stepDescription := translatedProcedureDescription(obj.ID, i, step, catalog.Translations)
		hasSanitizedPayloadExamples = hasSanitizedPayloadExamples || hasRawHTMLPayloadBlock(stepDescription)
		description := sanitizeMarkdown(stepDescription)
		procedure = append(procedure, map[string]string{
			"tactic":           step.Tactic,
			"tactic_name":      catalog.objectName(step.Tactic, tactic.Name),
			"technique":        step.Technique,
			"technique_name":   catalog.objectName(step.Technique, technique.Name),
			"description":      description,
			"description_line": sanitizeMarkdown(oneLine(stepDescription)),
		})
	}
	references := make([]map[string]string, 0, len(obj.References))
	for _, ref := range obj.References {
		references = append(references, map[string]string{"title": ref.Title, "url": safeExternalURL(ref.URL)})
	}
	params := map[string]any{
		"incident_date":             formatIncidentDate(obj.IncidentDate, obj.IncidentDateGranularity),
		"incident_date_raw":         obj.IncidentDate,
		"incident_date_granularity": obj.IncidentDateGranularity,
		"target":                    obj.Target,
		"actor":                     obj.Actor,
		"reporter":                  obj.Reporter,
		"case_study_type":           obj.CaseStudyType,
		"procedure_count":           len(obj.Procedure),
		"procedure":                 procedure,
		"references":                references,
	}
	if hasSanitizedPayloadExamples {
		params["has_sanitized_payload_examples"] = true
	}
	return Page{
		Section:     "studies",
		Kind:        "case-study",
		ID:          obj.ID,
		SourceName:  obj.Name,
		Title:       catalog.objectName(obj.ID, obj.Name),
		Description: body,
		Body:        b.String(),
		Params:      params,
	}
}

func (c Catalog) objectName(id, fallback string) string {
	if tr := strings.TrimSpace(c.Translations.Objects[id].Name); tr != "" {
		return tr
	}
	if fallback != "" {
		return fallback
	}
	if obj, ok := c.Tactics[id]; ok {
		return obj.Name
	}
	if obj, ok := c.Techniques[id]; ok {
		return obj.Name
	}
	if obj, ok := c.Mitigations[id]; ok {
		return obj.Name
	}
	if obj, ok := c.Studies[id]; ok {
		return obj.Name
	}
	return id
}

func (c Catalog) objectLink(section, id, fallback string) string {
	label := strings.TrimSpace(id + " " + c.objectName(id, fallback))
	return fmt.Sprintf("[%s](/%s/%s/)", escapeMarkdown(label), section, id)
}

func translatedName(id, source string, translations Translations) string {
	if tr := strings.TrimSpace(translations.Objects[id].Name); tr != "" {
		return tr
	}
	return source
}

func translatedBody(source, translated string) string {
	if tr := strings.TrimSpace(translated); tr != "" {
		return tr
	}
	return strings.TrimSpace(source)
}

func translatedProcedureDescription(studyID string, index int, step Procedure, translations Translations) string {
	tr := translations.Objects[studyID]
	candidate, ok := procedureTranslationForStep(tr, index, step)
	if !ok {
		return step.Description
	}
	if description := strings.TrimSpace(candidate.Description); description != "" {
		return description
	}
	return step.Description
}

func procedureTranslationForStep(tr Translation, index int, step Procedure) (ProcedureTranslation, bool) {
	if index >= 0 && index < len(tr.Procedure) {
		candidate := tr.Procedure[index]
		if procedureTranslationMatches(candidate, step) {
			return candidate, true
		}
	}
	for _, candidate := range tr.Procedure {
		if procedureTranslationMatches(candidate, step) {
			return candidate, true
		}
	}
	return ProcedureTranslation{}, false
}

func unusedProcedureTranslationForStep(tr Translation, index int, step Procedure, used map[int]bool) (int, ProcedureTranslation, bool) {
	if index >= 0 && index < len(tr.Procedure) && !used[index] {
		candidate := tr.Procedure[index]
		if procedureTranslationMatches(candidate, step) {
			return index, candidate, true
		}
	}
	for candidateIndex, candidate := range tr.Procedure {
		if used[candidateIndex] {
			continue
		}
		if procedureTranslationMatches(candidate, step) {
			return candidateIndex, candidate, true
		}
	}
	return -1, ProcedureTranslation{}, false
}

func translatedTechniqueUse(mitigationID string, technique TechniqueUse, translations Translations) string {
	for _, candidate := range translations.Objects[mitigationID].Techniques {
		if strings.TrimSpace(candidate.ID) != strings.TrimSpace(technique.ID) {
			continue
		}
		if use := strings.TrimSpace(candidate.Use); use != "" {
			return use
		}
		return technique.Use
	}
	if use := translatedTechniqueUseBySource(technique.Use, translations); use != "" {
		return use
	}
	return technique.Use
}

func translatedTechniqueUseBySource(source string, translations Translations) string {
	source = strings.TrimSpace(source)
	if source == "" {
		return ""
	}
	for _, candidate := range translations.TechniqueUses {
		if strings.TrimSpace(candidate.Source) != source {
			continue
		}
		return strings.TrimSpace(candidate.Use)
	}
	return ""
}

func procedureTranslationMatches(tr ProcedureTranslation, step Procedure) bool {
	return strings.TrimSpace(tr.Tactic) == strings.TrimSpace(step.Tactic) &&
		strings.TrimSpace(tr.Technique) == strings.TrimSpace(step.Technique)
}

func writeStatus(b *strings.Builder, translated string) {
	if strings.TrimSpace(translated) == "" {
		b.WriteString("> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.\n\n")
	}
}

func writeParagraphs(b *strings.Builder, text string) {
	for _, paragraph := range strings.Split(strings.TrimSpace(text), "\n\n") {
		paragraph = strings.TrimSpace(paragraph)
		if paragraph == "" {
			continue
		}
		b.WriteString(sanitizeMarkdown(paragraph))
		b.WriteString("\n\n")
	}
}

func writeMeta(b *strings.Builder, created, modified, maturity string) {
	if created == "" && modified == "" && maturity == "" {
		return
	}
	b.WriteString("\n## Метаданные\n\n")
	writeBullet(b, "Создано", created)
	writeBullet(b, "Обновлено", modified)
	writeBullet(b, "Зрелость", maturity)
}

func writeAttackReference(b *strings.Builder, ref *AttackReference) {
	if ref == nil || strings.TrimSpace(ref.ID) == "" {
		return
	}
	b.WriteString("\n## ATT&CK reference\n\n")
	if strings.TrimSpace(ref.URL) == "" {
		fmt.Fprintf(b, "- %s\n", ref.ID)
		return
	}
	refURL := safeExternalURL(ref.URL)
	if refURL == "" {
		fmt.Fprintf(b, "- %s\n", ref.ID)
		return
	}
	fmt.Fprintf(b, "- [%s](%s)\n", ref.ID, refURL)
}

func attackRefID(ref *AttackReference) string {
	if ref == nil {
		return ""
	}
	return strings.TrimSpace(ref.ID)
}

func attackRefURL(ref *AttackReference) string {
	if ref == nil {
		return ""
	}
	return safeExternalURL(ref.URL)
}

func writeTacticRelations(b *strings.Builder, title string, tacticIDs []string, catalog Catalog) {
	if len(tacticIDs) == 0 {
		return
	}
	b.WriteString("\n## ")
	b.WriteString(title)
	b.WriteString("\n\n<div class=\"relation-list\">\n")
	for _, tacticID := range tacticIDs {
		tactic := catalog.Tactics[tacticID]
		writeRelationItem(b, "tactics", tacticID, catalog.objectName(tacticID, tactic.Name), "", "")
	}
	b.WriteString("</div>\n\n")
}

func writeObjectRelations(b *strings.Builder, title, section string, objects []Object, catalog Catalog) {
	if len(objects) == 0 {
		return
	}
	b.WriteString("\n## ")
	b.WriteString(title)
	b.WriteString("\n\n<div class=\"relation-list\">\n")
	for _, obj := range objects {
		meta := ""
		if obj.SubtechniqueOf != "" {
			meta = "Подтехника"
		}
		writeRelationItem(b, section, obj.ID, catalog.objectName(obj.ID, obj.Name), meta, "")
	}
	b.WriteString("</div>\n\n")
}

func writeMitigationRelations(b *strings.Builder, entries []MitigationUseEntry, catalog Catalog) {
	if len(entries) == 0 {
		return
	}
	b.WriteString("\n## Меры защиты\n\n<div class=\"relation-list\">\n")
	for _, entry := range entries {
		writeRelationItem(
			b,
			"mitigations",
			entry.Mitigation.ID,
			catalog.objectName(entry.Mitigation.ID, entry.Mitigation.Name),
			"",
			translatedTechniqueUse(entry.Mitigation.ID, TechniqueUse{ID: entry.TechniqueID, Use: entry.Use}, catalog.Translations),
		)
	}
	b.WriteString("</div>\n\n")
}

func writeTechniqueUseRelations(b *strings.Builder, mitigationID string, techniques []TechniqueUse, catalog Catalog) {
	if len(techniques) == 0 {
		return
	}
	b.WriteString("\n## Связанные техники\n\n<div class=\"relation-list\">\n")
	for _, technique := range techniques {
		techniqueObj := catalog.Techniques[technique.ID]
		writeRelationItem(
			b,
			"techniques",
			technique.ID,
			catalog.objectName(technique.ID, techniqueObj.Name),
			"",
			translatedTechniqueUse(mitigationID, technique, catalog.Translations),
		)
	}
	b.WriteString("</div>\n\n")
}

func writeRelationItem(b *strings.Builder, section, id, name, meta, description string) {
	fmt.Fprintf(
		b,
		"<a class=\"relation-item\" href=\"/%s/%s/\"><span class=\"relation-id\">%s</span><strong>%s</strong>",
		html.EscapeString(section),
		html.EscapeString(id),
		html.EscapeString(id),
		html.EscapeString(name),
	)
	if strings.TrimSpace(meta) != "" {
		fmt.Fprintf(b, "<span class=\"relation-meta\">%s</span>", html.EscapeString(meta))
	}
	if strings.TrimSpace(description) != "" {
		b.WriteString(relationDescriptionHTML(description))
	}
	b.WriteString("</a>\n")
}

func writeProcedureExamples(b *strings.Builder, examples []ProcedureExample, catalog Catalog, limit int) {
	if len(examples) == 0 {
		return
	}
	b.WriteString("\n## Примеры процедур из кейсов\n\n<div class=\"relation-list procedure-relations\">\n")
	count := len(examples)
	if limit > 0 && count > limit {
		count = limit
	}
	for _, example := range examples[:count] {
		tactic := catalog.Tactics[example.Procedure.Tactic]
		meta := ""
		if strings.TrimSpace(example.Study.Actor) != "" {
			meta = "Актор: " + example.Study.Actor
		}
		if example.Procedure.Tactic != "" {
			if meta != "" {
				meta += " / "
			}
			meta += "Тактика: " + example.Procedure.Tactic + " " + catalog.objectName(example.Procedure.Tactic, tactic.Name)
		}
		writeRelationItem(
			b,
			"studies",
			example.Study.ID,
			catalog.objectName(example.Study.ID, example.Study.Name),
			meta,
			translatedProcedureDescription(example.Study.ID, example.Index, example.Procedure, catalog.Translations),
		)
	}
	b.WriteString("</div>\n\n")
	if limit > 0 && len(examples) > limit {
		fmt.Fprintf(b, "\nПоказано %d из %d примеров.\n", limit, len(examples))
	}
}

func writeReferences(b *strings.Builder, refs []Reference) {
	if len(refs) == 0 {
		return
	}
	b.WriteString("\n## Источники\n\n")
	for i, ref := range refs {
		title := strings.TrimSpace(ref.Title)
		if title == "" {
			title = strings.TrimSpace(ref.URL)
		}
		if title == "" {
			title = fmt.Sprintf("Источник %d", i+1)
		}
		refURL := safeExternalURL(ref.URL)
		if refURL == "" {
			fmt.Fprintf(b, "- %s\n", escapeMarkdown(title))
			continue
		}
		fmt.Fprintf(b, "- [%s](%s)\n", escapeMarkdown(title), refURL)
	}
}

func writeBullet(b *strings.Builder, label, value string) {
	if strings.TrimSpace(value) == "" {
		return
	}
	fmt.Fprintf(b, "- **%s:** %s\n", label, value)
}

func writeIDLinks(b *strings.Builder, title, section string, ids []string) {
	b.WriteString("\n## ")
	b.WriteString(title)
	b.WriteString("\n\n")
	for _, id := range ids {
		fmt.Fprintf(b, "- [%s](/%s/%s/)\n", id, section, id)
	}
}

func writePages(pages []Page) error {
	for _, section := range []string{"tactics", "techniques", "mitigations", "studies"} {
		if err := removeGeneratedPages(filepath.Join("content", section)); err != nil {
			return err
		}
	}
	for _, page := range pages {
		if err := writePage(page); err != nil {
			return err
		}
	}
	return nil
}

func removeGeneratedPages(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") || entry.Name() == "_index.md" {
			continue
		}
		path := filepath.Join(dir, entry.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if bytes.Contains(data, []byte("generated_by: "+generatedBy)) {
			if err := os.Remove(path); err != nil {
				return err
			}
		}
	}
	return nil
}

func writePage(page Page) error {
	if err := validateObjectID(page.ID); err != nil {
		return err
	}
	dir := filepath.Join("content", page.Section)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	frontMatter := map[string]any{
		"title":        page.Title,
		"atlas_id":     page.ID,
		"atlas_type":   page.Kind,
		"source_name":  page.SourceName,
		"url":          fmt.Sprintf("/%s/%s/", page.Section, page.ID),
		"generated":    true,
		"generated_by": generatedBy,
	}
	for key, value := range page.Params {
		frontMatter[key] = value
	}
	if strings.TrimSpace(page.Description) != "" {
		frontMatter["description"] = truncate(markdownPlainText(page.Description), 220)
	}
	encoded, err := yaml.Marshal(frontMatter)
	if err != nil {
		return err
	}
	path := filepath.Join(dir, page.ID+".md")
	body := strings.TrimRight(page.Body, "\n") + "\n"
	return os.WriteFile(path, []byte("---\n"+string(encoded)+"---\n\n"+body), 0o644)
}

func writeGeneratedAtlasData(atlas Atlas) error {
	if err := os.MkdirAll(filepath.Dir(generatedAtlasDataPath), 0o755); err != nil {
		return err
	}
	data, err := yaml.Marshal(atlas)
	if err != nil {
		return err
	}
	return os.WriteFile(generatedAtlasDataPath, data, 0o644)
}

func newCoverage(version string) *CoverageSummary {
	return &CoverageSummary{
		Version: version,
		Totals:  map[string]TypeSummary{},
	}
}

func (c *CoverageSummary) add(kind, id, sourceName, sourceField, sourceBody string, tr Translation) {
	nameNeedsReview, currentNameHash, translationNameHash := sourceNeedsReview(sourceName, translationSourceHash(tr, "name"))
	bodyNeedsReview, currentSourceHash, translationHash := sourceNeedsReview(sourceBody, translationSourceHash(tr, sourceField))
	hasBody := bodyField(tr, sourceField) != ""
	needsReview := nameNeedsReview || bodyNeedsReview
	item := CoverageItem{
		ID:          id,
		Name:        sourceName,
		Type:        kind,
		HasName:     strings.TrimSpace(tr.Name) != "",
		HasBody:     hasBody,
		SourceField: sourceField,
		NeedsReview: needsReview,
	}
	if nameNeedsReview {
		item.SourceField = "name"
		item.SourceSHA256 = currentNameHash
		item.TranslationSourceSHA256 = translationNameHash
		item.ReviewFields = append(item.ReviewFields, "name")
	}
	if bodyNeedsReview {
		if item.SourceSHA256 == "" {
			item.SourceField = sourceField
			item.SourceSHA256 = currentSourceHash
			item.TranslationSourceSHA256 = translationHash
		}
		item.ReviewFields = append(item.ReviewFields, sourceField)
	}
	if needsReview && len(item.ReviewFields) == 0 {
		item.SourceSHA256 = currentSourceHash
		item.TranslationSourceSHA256 = translationHash
	}
	total := c.Totals[kind]
	total.Total++
	if item.HasName {
		total.NameTranslated++
	}
	if item.HasBody {
		total.BodyTranslated++
	}
	if item.NeedsReview {
		total.NeedsReview++
	}
	if item.HasName && item.HasBody && !item.NeedsReview {
		total.FullyTranslated++
	}
	c.Totals[kind] = total

	if item.NeedsReview {
		c.Review = append(c.Review, item)
	}
	if !item.HasName && !item.HasBody {
		c.Missing = append(c.Missing, item)
		return
	}
	if !item.HasName || !item.HasBody {
		c.Partial = append(c.Partial, item)
	}
}

func (c *CoverageSummary) addCaseStudy(study Study, tr Translation) {
	nameNeedsReview, currentNameHash, translationNameHash := sourceNeedsReview(study.Name, translationSourceHash(tr, "name"))
	summaryNeedsReview, currentSourceHash, translationHash := sourceNeedsReview(study.Summary, translationSourceHash(tr, "summary"))
	hasBody := strings.TrimSpace(tr.Summary) != ""
	proceduresTotal, proceduresTranslated, procedureNeedsReview := procedureCoverage(study, tr)
	item := CoverageItem{
		ID:                   study.ID,
		Name:                 study.Name,
		Type:                 "case-study",
		HasName:              strings.TrimSpace(tr.Name) != "",
		HasBody:              hasBody,
		SourceField:          "summary",
		TotalProcedures:      proceduresTotal,
		TranslatedProcedures: proceduresTranslated,
		ProcedureNeedsReview: procedureNeedsReview,
		NeedsReview:          nameNeedsReview || summaryNeedsReview || procedureNeedsReview > 0,
	}
	if nameNeedsReview {
		item.SourceField = "name"
		item.SourceSHA256 = currentNameHash
		item.TranslationSourceSHA256 = translationNameHash
		item.ReviewFields = append(item.ReviewFields, "name")
	}
	if summaryNeedsReview {
		if item.SourceSHA256 == "" {
			item.SourceField = "summary"
			item.SourceSHA256 = currentSourceHash
			item.TranslationSourceSHA256 = translationHash
		}
		item.ReviewFields = append(item.ReviewFields, "summary")
	}
	if procedureNeedsReview > 0 {
		item.ReviewFields = append(item.ReviewFields, "procedure")
	}
	total := c.Totals[item.Type]
	total.Total++
	if item.HasName {
		total.NameTranslated++
	}
	if item.HasBody {
		total.BodyTranslated++
	}
	total.ProceduresTotal += proceduresTotal
	total.ProceduresTranslated += proceduresTranslated
	if item.NeedsReview {
		total.NeedsReview++
	}
	if item.HasName && item.HasBody && proceduresTranslated == proceduresTotal && !item.NeedsReview {
		total.FullyTranslated++
	}
	c.Totals[item.Type] = total

	if item.NeedsReview {
		c.Review = append(c.Review, item)
	}
	if !item.HasName && !item.HasBody && proceduresTranslated == 0 {
		c.Missing = append(c.Missing, item)
		return
	}
	if !item.HasName || !item.HasBody || proceduresTranslated < proceduresTotal {
		c.Partial = append(c.Partial, item)
	}
}

func (c *CoverageSummary) addResources(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".md" {
			continue
		}
		path := filepath.Join(dir, entry.Name())
		resource, body, err := readResourcePage(path)
		if err != nil {
			return err
		}
		id := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))
		if id == "_index" {
			id = "resources"
		}
		title := strings.TrimSpace(resource.Title)
		description := strings.TrimSpace(resource.Description)
		needsReview := hasEnglishProse(body)
		item := CoverageItem{
			ID:          id,
			Name:        title,
			Type:        "resource",
			HasName:     hasCyrillic(title),
			HasBody:     hasCyrillic(description),
			SourceField: "description",
			NeedsReview: needsReview,
		}
		total := c.Totals[item.Type]
		total.Total++
		if item.HasName {
			total.NameTranslated++
		}
		if item.HasBody {
			total.BodyTranslated++
		}
		if item.NeedsReview {
			total.NeedsReview++
		}
		if item.HasName && item.HasBody && !item.NeedsReview {
			total.FullyTranslated++
		}
		c.Totals[item.Type] = total

		if item.NeedsReview {
			c.Review = append(c.Review, item)
		}
		if !item.HasName && !item.HasBody {
			c.Missing = append(c.Missing, item)
			continue
		}
		if !item.HasName || !item.HasBody {
			c.Partial = append(c.Partial, item)
		}
	}
	return nil
}

type resourcePageFrontMatter struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	URL         string `yaml:"url"`
}

func readResourcePage(path string) (resourcePageFrontMatter, string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return resourcePageFrontMatter{}, "", err
	}
	text := string(data)
	if !strings.HasPrefix(text, "---\n") {
		return resourcePageFrontMatter{}, text, nil
	}
	rest := strings.TrimPrefix(text, "---\n")
	parts := strings.SplitN(rest, "\n---\n", 2)
	if len(parts) != 2 {
		return resourcePageFrontMatter{}, "", fmt.Errorf("%s: invalid front matter", path)
	}
	var resource resourcePageFrontMatter
	if err := yaml.Unmarshal([]byte(parts[0]), &resource); err != nil {
		return resourcePageFrontMatter{}, "", fmt.Errorf("%s: %w", path, err)
	}
	return resource, parts[1], nil
}

func hasCyrillic(value string) bool {
	return cyrillicPattern.MatchString(value)
}

func hasEnglishProse(markdown string) bool {
	text := stripMarkdownNoise(markdown)
	for _, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if !hasCyrillic(line) && englishProsePattern.MatchString(line) {
			return true
		}
		if hasCyrillic(line) && longEnglishProsePattern.MatchString(line) {
			return true
		}
	}
	return false
}

func procedureCoverage(study Study, tr Translation) (total int, translated int, needsReview int) {
	total = len(study.Procedure)
	usedTranslations := map[int]bool{}
	for index, step := range study.Procedure {
		candidateIndex, candidate, ok := unusedProcedureTranslationForStep(tr, index, step, usedTranslations)
		if !ok || strings.TrimSpace(candidate.Description) == "" {
			continue
		}
		usedTranslations[candidateIndex] = true
		translated++
		translationHash := strings.TrimSpace(candidate.Source.DescriptionSHA256)
		if translationHash != "" && translationHash != sourceHash(step.Description) {
			needsReview++
		}
	}
	return total, translated, needsReview
}

func sourceHash(value string) string {
	normalized := strings.TrimSpace(value)
	if normalized == "" {
		return ""
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(normalized)))
}

func sourceNeedsReview(sourceValue, translationHash string) (bool, string, string) {
	currentSourceHash := sourceHash(sourceValue)
	translationHash = strings.TrimSpace(translationHash)
	return translationHash != "" && currentSourceHash != "" && translationHash != currentSourceHash, currentSourceHash, translationHash
}

func translationSourceHash(tr Translation, field string) string {
	switch field {
	case "name":
		return strings.TrimSpace(tr.Source.NameSHA256)
	case "summary":
		return strings.TrimSpace(tr.Source.SummarySHA256)
	default:
		return strings.TrimSpace(tr.Source.DescriptionSHA256)
	}
}

func bodyField(tr Translation, field string) string {
	switch field {
	case "summary":
		return strings.TrimSpace(tr.Summary)
	default:
		return strings.TrimSpace(tr.Description)
	}
}

func writeCoverageReports(coverage *CoverageSummary) error {
	sort.Slice(coverage.Missing, func(i, j int) bool {
		return coverage.Missing[i].ID < coverage.Missing[j].ID
	})
	sort.Slice(coverage.Partial, func(i, j int) bool {
		return coverage.Partial[i].ID < coverage.Partial[j].ID
	})
	sort.Slice(coverage.Review, func(i, j int) bool {
		return coverage.Review[i].ID < coverage.Review[j].ID
	})

	if err := os.MkdirAll(filepath.Dir(reportDataPath), 0o755); err != nil {
		return err
	}
	data, err := yaml.Marshal(coverage)
	if err != nil {
		return err
	}
	if err := os.WriteFile(reportDataPath, data, 0o644); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(reportPath), 0o755); err != nil {
		return err
	}
	var b strings.Builder
	fmt.Fprintf(&b, "Версия данных ATLAS: `%s`\n\n", coverage.Version)
	b.WriteString("| Тип объектов | Всего | Названий переведено | Описаний переведено | Процедур переведено | Полностью переведено | Требует проверки |\n")
	b.WriteString("| --- | ---: | ---: | ---: | ---: | ---: | ---: |\n")
	for _, kind := range sortedKinds(coverage.Totals) {
		total := coverage.Totals[kind]
		fmt.Fprintf(&b, "| %s | %d | %d | %d | %s | %d | %d |\n", coverageKindLabel(kind), total.Total, total.NameTranslated, total.BodyTranslated, procedureSummary(total), total.FullyTranslated, total.NeedsReview)
	}
	fmt.Fprintf(&b, "\n## Требует проверки (%d)\n\n", len(coverage.Review))
	writeCoverageList(&b, coverage.Review)
	fmt.Fprintf(&b, "\n## Частичный перевод (%d)\n\n", len(coverage.Partial))
	writeCoverageList(&b, coverage.Partial)
	fmt.Fprintf(&b, "\n## Перевод отсутствует (%d)\n\n", len(coverage.Missing))
	writeCoverageList(&b, coverage.Missing)

	report := []byte("# Покрытие перевода\n\n" + b.String())
	if err := os.WriteFile(reportPath, report, 0o644); err != nil {
		return err
	}
	var pageBody strings.Builder
	fmt.Fprintf(&pageBody, "Версия данных ATLAS: `%s`\n\n", coverage.Version)
	writeCoverageHTMLTable(&pageBody, coverage)
	fmt.Fprintf(&pageBody, "\n## Требует проверки (%d)\n\n", len(coverage.Review))
	writeCoverageList(&pageBody, coverage.Review)
	fmt.Fprintf(&pageBody, "\n## Частичный перевод (%d)\n\n", len(coverage.Partial))
	writeCoverageList(&pageBody, coverage.Partial)
	fmt.Fprintf(&pageBody, "\n## Перевод отсутствует (%d)\n\n", len(coverage.Missing))
	writeCoverageList(&pageBody, coverage.Missing)
	page := "---\ntitle: \"Покрытие перевода\"\nurl: \"/translation-coverage/\"\ngenerated: true\ngenerated_by: atlasgen\n---\n\n" + pageBody.String()
	return os.WriteFile(reportPagePath, []byte(page), 0o644)
}

func writeCoverageHTMLTable(b *strings.Builder, coverage *CoverageSummary) {
	columns := []string{
		"Тип объектов",
		"Всего",
		"Названий переведено",
		"Описаний переведено",
		"Процедур переведено",
		"Полностью переведено",
		"Требует проверки",
	}
	b.WriteString(`<table class="coverage-table coverage-detail-table">` + "\n<thead>\n<tr>\n")
	for _, column := range columns {
		fmt.Fprintf(b, "<th>%s</th>\n", html.EscapeString(column))
	}
	b.WriteString("</tr>\n</thead>\n<tbody>\n")
	for _, kind := range sortedKinds(coverage.Totals) {
		total := coverage.Totals[kind]
		values := []string{
			coverageKindLabel(kind),
			fmt.Sprintf("%d", total.Total),
			fmt.Sprintf("%d", total.NameTranslated),
			fmt.Sprintf("%d", total.BodyTranslated),
			procedureSummary(total),
			fmt.Sprintf("%d", total.FullyTranslated),
			fmt.Sprintf("%d", total.NeedsReview),
		}
		b.WriteString("<tr>\n")
		for i, value := range values {
			if i == 0 {
				fmt.Fprintf(b, `<th scope="row" data-label="%s">%s</th>`+"\n", html.EscapeString(columns[i]), html.EscapeString(value))
				continue
			}
			fmt.Fprintf(b, `<td data-label="%s">%s</td>`+"\n", html.EscapeString(columns[i]), html.EscapeString(value))
		}
		b.WriteString("</tr>\n")
	}
	b.WriteString("</tbody>\n</table>\n")
}

func sortedKinds(totals map[string]TypeSummary) []string {
	kinds := make([]string, 0, len(totals))
	for kind := range totals {
		kinds = append(kinds, kind)
	}
	sort.Strings(kinds)
	return kinds
}

func writeCoverageList(b *strings.Builder, items []CoverageItem) {
	if len(items) == 0 {
		b.WriteString("Нет.\n")
		return
	}
	for _, item := range items {
		pageURL := coverageItemURL(item)
		idLabel := fmt.Sprintf("`%s`", item.ID)
		if pageURL != "" {
			idLabel = fmt.Sprintf("[`%s`](%s)", item.ID, pageURL)
		}
		fmt.Fprintf(
			b,
			"- %s (%s): %s; название: %s; описание/summary: %s%s%s\n",
			idLabel,
			coverageKindLabel(item.Type),
			item.Name,
			yesNo(item.HasName),
			yesNo(item.HasBody),
			procedureLabel(item),
			reviewLabel(item),
		)
	}
}

func procedureSummary(total TypeSummary) string {
	if total.ProceduresTotal == 0 {
		return "-"
	}
	return fmt.Sprintf("%d/%d", total.ProceduresTranslated, total.ProceduresTotal)
}

func procedureLabel(item CoverageItem) string {
	if item.TotalProcedures == 0 {
		return ""
	}
	label := fmt.Sprintf("; процедуры: %d/%d", item.TranslatedProcedures, item.TotalProcedures)
	if item.ProcedureNeedsReview > 0 {
		label += fmt.Sprintf("; процедуры требуют проверки: %d", item.ProcedureNeedsReview)
	}
	return label
}

func reviewLabel(item CoverageItem) string {
	if !item.NeedsReview {
		return ""
	}
	if len(item.ReviewFields) == 0 {
		return "; требует проверки: да"
	}
	labels := make([]string, 0, len(item.ReviewFields))
	for _, field := range item.ReviewFields {
		labels = append(labels, reviewFieldLabel(field))
	}
	return "; требует проверки: " + strings.Join(labels, ", ")
}

func reviewFieldLabel(field string) string {
	switch field {
	case "name":
		return "название"
	case "description":
		return "описание"
	case "summary":
		return "summary"
	case "procedure":
		return "процедуры"
	default:
		return field
	}
}

func coverageItemURL(item CoverageItem) string {
	switch item.Type {
	case "case-study":
		return "/studies/" + item.ID + "/"
	case "mitigation":
		return "/mitigations/" + item.ID + "/"
	case "resource":
		if item.ID == "resources" {
			return "/resources/"
		}
		return "/resources/" + item.ID + "/"
	case "tactic":
		return "/tactics/" + item.ID + "/"
	case "technique":
		return "/techniques/" + item.ID + "/"
	default:
		return ""
	}
}

func coverageKindLabel(kind string) string {
	switch kind {
	case "case-study":
		return "Кейсы"
	case "mitigation":
		return "Меры защиты"
	case "resource":
		return "Ресурсы"
	case "tactic":
		return "Тактики"
	case "technique":
		return "Техники"
	default:
		return kind
	}
}

func yesNo(value bool) string {
	if value {
		return "да"
	}
	return "нет"
}

func sortObjects(objects []Object) {
	sort.Slice(objects, func(i, j int) bool {
		return objects[i].ID < objects[j].ID
	})
}

func sortProcedureExamples(examples []ProcedureExample) {
	sort.Slice(examples, func(i, j int) bool {
		if examples[i].Study.ID == examples[j].Study.ID {
			if examples[i].Procedure.Tactic == examples[j].Procedure.Tactic {
				return examples[i].Procedure.Technique < examples[j].Procedure.Technique
			}
			return examples[i].Procedure.Tactic < examples[j].Procedure.Tactic
		}
		return examples[i].Study.ID < examples[j].Study.ID
	})
}

func formatIncidentDate(date, granularity string) string {
	date = strings.TrimSpace(date)
	if date == "" {
		return ""
	}
	parts := strings.Split(date, "-")
	switch strings.ToUpper(strings.TrimSpace(granularity)) {
	case "YEAR":
		if len(parts) >= 1 {
			return parts[0]
		}
	case "MONTH":
		if len(parts) >= 2 {
			return parts[0] + "-" + parts[1]
		}
	case "DATE", "DAY":
		return date
	}
	if granularity == "" {
		return date
	}
	return fmt.Sprintf("%s (%s)", date, granularity)
}

func escapeMarkdown(value string) string {
	replacer := strings.NewReplacer("[", "\\[", "]", "\\]")
	return replacer.Replace(value)
}

func oneLine(value string) string {
	return strings.Join(strings.Fields(value), " ")
}

func relationDescriptionSummary(value string) string {
	lines := strings.Split(strings.TrimSpace(value), "\n")
	visibleLines := make([]string, 0, len(lines))
	inFence := false
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if isFenceDelimiter(line) {
			inFence = !inFence
			continue
		}
		if inFence {
			continue
		}
		if isMarkdownTableLine(line) {
			tableLines := []string{line}
			for i+1 < len(lines) && isMarkdownTableLine(strings.TrimSpace(lines[i+1])) {
				i++
				tableLines = append(tableLines, strings.TrimSpace(lines[i]))
			}
			if summary := summarizeMarkdownTable(tableLines); summary != "" {
				visibleLines = append(visibleLines, summary)
			}
			continue
		}
		visibleLines = append(visibleLines, line)
	}
	return removeDanglingBlockIntro(oneLine(strings.Join(visibleLines, "\n")))
}

func relationDescriptionHTML(value string) string {
	if listHTML := relationDescriptionListHTML(value); listHTML != "" {
		return listHTML
	}
	if summary := relationDescriptionSummary(value); summary != "" {
		return fmt.Sprintf("<p>%s</p>", html.EscapeString(summary))
	}
	return ""
}

func relationDescriptionListHTML(value string) string {
	lines := visibleNonFenceLines(value)
	title := ""
	items := []string{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "- ") {
			if title == "" && len(items) == 0 {
				return ""
			}
			items = append(items, strings.TrimSpace(strings.TrimPrefix(line, "- ")))
			continue
		}
		if title != "" || len(items) > 0 {
			return ""
		}
		title = line
	}
	if len(items) == 0 {
		return ""
	}
	var b strings.Builder
	if title != "" {
		fmt.Fprintf(&b, "<p>%s</p>", html.EscapeString(title))
	}
	b.WriteString("<ul>")
	for _, item := range items {
		fmt.Fprintf(&b, "<li>%s</li>", html.EscapeString(item))
	}
	b.WriteString("</ul>")
	return b.String()
}

func visibleNonFenceLines(value string) []string {
	lines := strings.Split(strings.TrimSpace(value), "\n")
	visibleLines := make([]string, 0, len(lines))
	inFence := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if isFenceDelimiter(trimmed) {
			inFence = !inFence
			continue
		}
		if inFence {
			continue
		}
		visibleLines = append(visibleLines, line)
	}
	return visibleLines
}

func isMarkdownTableLine(value string) bool {
	value = strings.TrimSpace(value)
	return strings.HasPrefix(value, "|") && strings.HasSuffix(value, "|") && strings.Count(value, "|") >= 2
}

func summarizeMarkdownTable(lines []string) string {
	entries := make([]string, 0, len(lines))
	headerSeen := false
	for _, line := range lines {
		cells := splitMarkdownTableCells(line)
		if len(cells) < 2 || isMarkdownTableSeparator(cells) {
			continue
		}
		if !headerSeen {
			headerSeen = true
			continue
		}
		entries = append(entries, cells[0]+" - "+cells[1])
	}
	return strings.Join(entries, "; ")
}

func splitMarkdownTableCells(line string) []string {
	line = strings.Trim(strings.TrimSpace(line), "|")
	parts := strings.Split(line, "|")
	cells := make([]string, 0, len(parts))
	for _, part := range parts {
		cell := strings.TrimSpace(part)
		if cell != "" {
			cells = append(cells, cell)
		}
	}
	return cells
}

func isMarkdownTableSeparator(cells []string) bool {
	if len(cells) == 0 {
		return false
	}
	for _, cell := range cells {
		if strings.Trim(cell, " :-") != "" {
			return false
		}
	}
	return true
}

var (
	markdownLinkPattern          = regexp.MustCompile(`!?\[([^\]]*)\]\([^)]+\)`)
	markdownInlineLinkPattern    = regexp.MustCompile(`!?\[([^\]]*)\]\(([^)\s]+)(?:\s+[^)]*)?\)`)
	markdownReferencePattern     = regexp.MustCompile(`(?m)^\[[^\]]+\]:\s+\S+.*$`)
	markdownReferenceLinkPattern = regexp.MustCompile(`(?m)^(\[[^\]]+\]:\s+)(\S+)(.*)$`)
	markdownReferenceUsePattern  = regexp.MustCompile(`!?\[([^\]]*)\]\[[^\]]*\]`)
	markdownMarkupPattern        = regexp.MustCompile("[`*_#>]")
	htmlTagPattern               = regexp.MustCompile(`<[^>]+>`)
	cyrillicPattern              = regexp.MustCompile(`[А-Яа-яЁё]`)
	englishProsePattern          = regexp.MustCompile(`[A-Za-z]{2,}(?:[ \t]+[A-Za-z]{2,}){3,}`)
	longEnglishProsePattern      = regexp.MustCompile(`[A-Za-z]{2,}(?:[ \t]+[A-Za-z]{2,}){7,}`)
	objectIDPattern              = regexp.MustCompile(`^[A-Za-z0-9]+(?:\.[A-Za-z0-9]+){1,2}$`)
	relationBlockIntroPatterns   = []*regexp.Regexp{
		regexp.MustCompile(`\s+(?:[Нн]ачальный\s+)?[Фф]рагмент.*:\s*$`),
		regexp.MustCompile(`\s+Видимая строка и скрытый фрагмент:\s*$`),
		regexp.MustCompile(`\s+Текст промпта[^.!?]*:\s*$`),
		regexp.MustCompile(`\s+Для этого[^.!?]*(?:полезная нагрузка|payload)[^.!?]*:\s*$`),
	}
)

func removeDanglingBlockIntro(value string) string {
	value = strings.TrimSpace(value)
	for _, pattern := range relationBlockIntroPatterns {
		value = pattern.ReplaceAllString(value, "")
	}
	return strings.TrimSpace(value)
}

func sanitizeMarkdown(value string) string {
	value = fenceRawHTMLBlocks(value)
	lines := strings.Split(value, "\n")
	inFence := false
	for i, line := range lines {
		if isFenceDelimiter(line) {
			inFence = !inFence
			continue
		}
		if inFence {
			continue
		}
		lines[i] = sanitizeMarkdownLine(line)
	}
	return strings.Join(lines, "\n")
}

func hasRawHTMLPayloadBlock(value string) bool {
	inFence := false
	for _, line := range strings.Split(value, "\n") {
		if isFenceDelimiter(line) {
			inFence = !inFence
			continue
		}
		if !inFence && isRawHTMLPayloadLine(line) {
			return true
		}
	}
	return false
}

func fenceRawHTMLBlocks(value string) string {
	lines := strings.Split(value, "\n")
	result := make([]string, 0, len(lines)+4)
	inGeneratedFence := false
	generatedFenceUntilDivClose := false
	inExistingFence := false
	for _, line := range lines {
		if isFenceDelimiter(line) {
			if inGeneratedFence {
				result = append(result, "```")
				inGeneratedFence = false
				generatedFenceUntilDivClose = false
			}
			inExistingFence = !inExistingFence
			result = append(result, line)
			continue
		}
		if inExistingFence {
			result = append(result, line)
			continue
		}
		if inGeneratedFence && generatedFenceUntilDivClose {
			result = append(result, line)
			if strings.Contains(strings.ToLower(line), "</div") {
				result = append(result, "```")
				result = append(result, "")
				inGeneratedFence = false
				generatedFenceUntilDivClose = false
			}
			continue
		}
		if isRawHTMLPayloadLine(line) {
			if !inGeneratedFence {
				if len(result) > 0 && strings.TrimSpace(result[len(result)-1]) != "" {
					result = append(result, "")
				}
				result = append(result, "```html")
				inGeneratedFence = true
				generatedFenceUntilDivClose = strings.HasPrefix(strings.ToLower(strings.TrimSpace(line)), "<div") &&
					!strings.Contains(strings.ToLower(line), "</div")
			}
			result = append(result, line)
			if generatedFenceUntilDivClose && strings.Contains(strings.ToLower(line), "</div") {
				result = append(result, "```")
				result = append(result, "")
				inGeneratedFence = false
				generatedFenceUntilDivClose = false
			}
			continue
		}
		if inGeneratedFence {
			result = append(result, "```")
			result = append(result, "")
			inGeneratedFence = false
			generatedFenceUntilDivClose = false
		}
		result = append(result, line)
	}
	if inGeneratedFence {
		result = append(result, "```")
	}
	return strings.Join(result, "\n")
}

func isRawHTMLPayloadLine(value string) bool {
	value = strings.TrimSpace(value)
	return strings.HasPrefix(value, "<div") ||
		strings.HasPrefix(value, "</div") ||
		strings.HasPrefix(value, "<span") ||
		strings.HasPrefix(value, "</span")
}

func sanitizeMarkdownLine(value string) string {
	value = markdownReferenceLinkPattern.ReplaceAllStringFunc(value, func(match string) string {
		parts := markdownReferenceLinkPattern.FindStringSubmatch(match)
		if len(parts) != 4 {
			return match
		}
		destination := strings.TrimSpace(parts[2])
		if !isSafeMarkdownURL(destination) {
			return ""
		}
		return parts[1] + destination + parts[3]
	})
	value = markdownInlineLinkPattern.ReplaceAllStringFunc(value, func(match string) string {
		parts := markdownInlineLinkPattern.FindStringSubmatch(match)
		if len(parts) != 3 {
			return match
		}
		label := escapeMarkdown(parts[1])
		destination := strings.TrimSpace(parts[2])
		if !isSafeMarkdownURL(destination) {
			return label
		}
		if strings.HasPrefix(match, "![") {
			return fmt.Sprintf("![%s](%s)", label, destination)
		}
		return fmt.Sprintf("[%s](%s)", label, destination)
	})
	placeholders := []struct {
		tag   string
		token string
	}{
		{tag: "<sup>", token: "\x00ATLAS_SUP_OPEN\x00"},
		{tag: "</sup>", token: "\x00ATLAS_SUP_CLOSE\x00"},
		{tag: "<br>", token: "\x00ATLAS_BR\x00"},
		{tag: "<br/>", token: "\x00ATLAS_BR_SLASH\x00"},
		{tag: "<br />", token: "\x00ATLAS_BR_SPACE\x00"},
	}
	for _, placeholder := range placeholders {
		value = strings.ReplaceAll(value, placeholder.tag, placeholder.token)
	}
	replacer := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	value = replacer.Replace(value)
	for _, placeholder := range placeholders {
		value = strings.ReplaceAll(value, placeholder.token, placeholder.tag)
	}
	return value
}

func isFenceDelimiter(value string) bool {
	value = strings.TrimSpace(value)
	return strings.HasPrefix(value, "```") || strings.HasPrefix(value, "~~~")
}

func safeExternalURL(value string) string {
	value = strings.TrimSpace(value)
	if value == "" || !isSafeAbsoluteURL(value) {
		return ""
	}
	return value
}

func isSafeMarkdownURL(value string) bool {
	value = strings.TrimSpace(value)
	if strings.HasPrefix(value, "/") && !strings.HasPrefix(value, "//") && !hasControlCharacter(value) {
		return true
	}
	return isSafeAbsoluteURL(value)
}

func isSafeAbsoluteURL(value string) bool {
	if hasControlCharacter(value) {
		return false
	}
	parsed, err := url.Parse(value)
	if err != nil || parsed.Host == "" {
		return false
	}
	scheme := strings.ToLower(parsed.Scheme)
	return scheme == "http" || scheme == "https"
}

func hasControlCharacter(value string) bool {
	for _, r := range value {
		if r < 0x20 || r == 0x7f {
			return true
		}
	}
	return false
}

func validateObjectID(id string) error {
	if objectIDPattern.MatchString(id) {
		return nil
	}
	return fmt.Errorf("unsafe ATLAS object id %q", id)
}

func markdownPlainText(value string) string {
	value = markdownReferencePattern.ReplaceAllString(value, "")
	value = markdownLinkPattern.ReplaceAllString(value, "$1")
	value = markdownMarkupPattern.ReplaceAllString(value, "")
	return oneLine(value)
}

func stripMarkdownNoise(value string) string {
	value = markdownReferencePattern.ReplaceAllString(value, "")
	value = markdownInlineLinkPattern.ReplaceAllStringFunc(value, func(match string) string {
		parts := markdownInlineLinkPattern.FindStringSubmatch(match)
		if len(parts) < 3 {
			return ""
		}
		linkText := parts[1]
		target := parts[2]
		if strings.HasPrefix(target, "http://") || strings.HasPrefix(target, "https://") || strings.HasPrefix(target, "mailto:") {
			if hasCyrillic(linkText) {
				return linkText
			}
			return " "
		}
		return "\n" + linkText + "\n"
	})
	value = markdownReferenceUsePattern.ReplaceAllString(value, " ")
	value = markdownLinkPattern.ReplaceAllString(value, "$1")
	value = htmlTagPattern.ReplaceAllString(value, " ")
	value = markdownMarkupPattern.ReplaceAllString(value, "")
	return value
}

func truncate(value string, limit int) string {
	runes := []rune(value)
	if len(runes) <= limit {
		return value
	}
	if limit <= 1 {
		return string(runes[:limit])
	}
	candidate := strings.TrimSpace(string(runes[:limit-1]))
	if lastSpace := strings.LastIndex(candidate, " "); lastSpace > limit/2 {
		candidate = strings.TrimSpace(candidate[:lastSpace])
	}
	return candidate + "..."
}

---
actor: DeepSeek, Moonshot AI, MiniMax
atlas_id: AML.CS0056
atlas_type: case-study
case_study_type: incident
description: 'Anthropic выявила кампании по извлечению возможностей Claude, проводившиеся тремя китайскими ИИ-лабораториями: DeepSeek, Moonshot и MiniMax. В совокупности в этих кампаниях использовалось примерно 24 000 учетных...'
generated: true
generated_by: atlasgen
incident_date: "2026-02-23"
incident_date_granularity: Day
incident_date_raw: "2026-02-23"
procedure:
    - description: DeepSeek, Moonshot AI и MiniMax использовали коммерческие прокси-сервисы для доступа к Claude, обходя политику Anthropic, запрещающую коммерческий доступ к Claude в Китае.
      description_line: DeepSeek, Moonshot AI и MiniMax использовали коммерческие прокси-сервисы для доступа к Claude, обходя политику Anthropic, запрещающую коммерческий доступ к Claude в Китае.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0008.005
      technique_name: Прокси для ИИ-сервисов
    - description: DeepSeek, Moonshot AI и MiniMax сгенерировали крупные наборы промптов для извлечения возможностей Claude.
      description_line: DeepSeek, Moonshot AI и MiniMax сгенерировали крупные наборы промптов для извлечения возможностей Claude.
      tactic: AML.TA0001
      tactic_name: Адаптация атак, связанных с ИИ
      technique: AML.T0065
      technique_name: Создание промптов для LLM
    - description: ИИ-лаборатории обращались к API инференса Claude через примерно 24 000 поддельных учетных записей в совокупности.
      description_line: ИИ-лаборатории обращались к API инференса Claude через примерно 24 000 поддельных учетных записей в совокупности.
      tactic: AML.TA0000
      tactic_name: Доступ к ИИ-модели
      technique: AML.T0040
      technique_name: Доступ к API инференса ИИ-модели
    - description: DeepSeek, Moonshot AI и MiniMax использовали сгенерированные промпты для многократных запросов к Claude и обучения собственных моделей на его ответах. В совокупности в ходе кампаний дистилляции лаборатории отправили более 16 млн запросов.
      description_line: DeepSeek, Moonshot AI и MiniMax использовали сгенерированные промпты для многократных запросов к Claude и обучения собственных моделей на его ответах. В совокупности в ходе кампаний дистилляции лаборатории отправили более 16 млн запросов.
      tactic: AML.TA0010
      tactic_name: Эксфильтрация
      technique: AML.T0024.002
      technique_name: Извлечение ИИ-модели
    - description: DeepSeek, Moonshot AI и MiniMax переняли возможности Claude с помощью дистилляции, затратив лишь малую часть стоимости разработки собственных моделей. Их интересовали наиболее отличительные возможности Claude, включая агентное рассуждение, использование инструментов и генерацию кода.
      description_line: DeepSeek, Moonshot AI и MiniMax переняли возможности Claude с помощью дистилляции, затратив лишь малую часть стоимости разработки собственных моделей. Их интересовали наиболее отличительные возможности Claude, включая агентное рассуждение, использование инструментов и генерацию кода.
      tactic: AML.TA0011
      tactic_name: Воздействие
      technique: AML.T0048.004
      technique_name: Кража интеллектуальной собственности ИИ
    - description: 'У дистиллированных моделей отсутствуют защитные механизмы, поэтому их можно использовать во вредоносных целях: для наступательных киберопераций, кампаний дезинформации, массовой слежки и цензуры.'
      description_line: 'У дистиллированных моделей отсутствуют защитные механизмы, поэтому их можно использовать во вредоносных целях: для наступательных киберопераций, кампаний дезинформации, массовой слежки и цензуры.'
      tactic: AML.TA0011
      tactic_name: Воздействие
      technique: AML.T0048.002
      technique_name: Общественный вред
    - description: У дистиллированных моделей нет гардрейлов безопасности Claude, в результате чего пользователи могут столкнуться с вредоносными ответами и поведением модели.
      description_line: У дистиллированных моделей нет гардрейлов безопасности Claude, в результате чего пользователи могут столкнуться с вредоносными ответами и поведением модели.
      tactic: AML.TA0011
      tactic_name: Воздействие
      technique: AML.T0048.003
      technique_name: Ущерб пользователям
procedure_count: 7
references:
    - title: Detecting and preventing distillation attacks
      url: https://www.anthropic.com/news/detecting-and-preventing-distillation-attacks
reporter: Anthropic
source_name: Model Distillation Campaigns Targeting Anthropic Claude
target: Anthropic Claude
title: Кампании по дистилляции моделей, нацеленные на Anthropic Claude
url: /studies/AML.CS0056/
---

Anthropic выявила кампании по извлечению возможностей Claude, проводившиеся тремя китайскими ИИ-лабораториями: DeepSeek, Moonshot и MiniMax. В совокупности в этих кампаниях использовалось примерно 24 000 учетных записей и 16 млн запросов. Они применяли дистилляцию моделей, чтобы обучать собственные модели на ответах Claude в попытке воспроизвести возможности Claude, включая агентное рассуждение, генерацию кода, использование инструментов и работу с компьютером.

Как указано в отчете Anthropic, эти лаборатории использовали дистилляцию моделей как способ обойти экспортные ограничения Anthropic.[1] Дистиллированные модели лишены защитных механизмов, которые не позволяют злоумышленникам использовать передовые модели во вредоносных целях, таких как разработка биооружия, дезинформация, наступательные кибероперации и массовая слежка.

[1]: https://www.anthropic.com/news/detecting-and-preventing-distillation-attacks "Detecting and preventing distillation attacks"

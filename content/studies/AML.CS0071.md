---
actor: Unknown Chinese-language actor
atlas_id: AML.CS0071
atlas_type: case-study
case_study_type: incident
description: In early July 2026, an unknown Chinese-language operator used a multi-agent framework built on Hermes and OpenClaw against government systems that subsequent public reporting identified as...
generated: true
generated_by: atlasgen
incident_date: "2026-07-01"
incident_date_granularity: Day
incident_date_raw: "2026-07-01"
procedure:
    - description: The framework dispatched up to eight specialized agents concurrently across 12 attack waves. It assigned separate reconnaissance, authentication, API-testing, vulnerability-research, credential, and exploitation missions; allocated additional testing to promising findings; requested independent validation; aggregated after-action reports; and redirected subsequent work based on the status of related workstreams.
      description_line: The framework dispatched up to eight specialized agents concurrently across 12 attack waves. It assigned separate reconnaissance, authentication, API-testing, vulnerability-research, credential, and exploitation missions; allocated additional testing to promising findings; requested independent validation; aggregated after-action reports; and redirected subsequent work based on the status of related workstreams.
      tactic: AML.TA0001
      tactic_name: Подготовка атаки на ИИ
      technique: AML.T0124
      technique_name: Autonomous Attack Orchestration
    - description: The framework constructed numerous candidate multi-step attack paths using confirmed prerequisites, observed blockers, and estimated success probabilities. It promoted paths supported by validated evidence, queued paths requiring additional investigation, discarded false positives and blocked paths, and initiated target-specific learning cycles when existing methods failed.
      description_line: The framework constructed numerous candidate multi-step attack paths using confirmed prerequisites, observed blockers, and estimated success probabilities. It promoted paths supported by validated evidence, queued paths requiring additional investigation, discarded false positives and blocked paths, and initiated target-specific learning cycles when existing methods failed.
      tactic: AML.TA0001
      tactic_name: Подготовка атаки на ИИ
      technique: AML.T0117
      technique_name: Автономная адаптация пути атаки
    - description: The framework exchanged assignments, findings, validation results, status, and after-action information between its orchestrating control process and specialized sub-agents. Aggregated results informed later assignments and attack waves.
      description_line: The framework exchanged assignments, findings, validation results, status, and after-action information between its orchestrating control process and specialized sub-agents. Aggregated results informed later assignments and attack waves.
      tactic: AML.TA0001
      tactic_name: Подготовка атаки на ИИ
      technique: AML.T0118.001
      technique_name: Direct Agent Communication
    - description: The framework performed reconnaissance on an internet-facing Taiwanese government portal, interpreting client-side application bundles, following discovered infrastructure relationships, and generating additional reconnaissance objectives. It identified 21 connected systems, six SSO sub-realms, authentication configuration, signing-key information, and more than 36 API endpoints on one system.
      description_line: The framework performed reconnaissance on an internet-facing Taiwanese government portal, interpreting client-side application bundles, following discovered infrastructure relationships, and generating additional reconnaissance objectives. It identified 21 connected systems, six SSO sub-realms, authentication configuration, signing-key information, and more than 36 API endpoints on one system.
      tactic: AML.TA0002
      tactic_name: Разведка
      technique: AML.T0116
      technique_name: Автономная разведка
    - description: The framework probed primary government applications and APIs for exposed interfaces, authentication behavior, misconfigurations, and vulnerabilities. This scanning identified multiple potential paths into the targeted systems.
      description_line: The framework probed primary government applications and APIs for exposed interfaces, authentication behavior, misconfigurations, and vulnerabilities. This scanning identified multiple potential paths into the targeted systems.
      tactic: AML.TA0002
      tactic_name: Разведка
      technique: AML.T0006
      technique_name: Активное сканирование
    - description: The framework retrieved employee names, departments, identifiers, and SSO account information from an exposed user-database API without authentication.
      description_line: The framework retrieved employee names, departments, identifiers, and SSO account information from an exposed user-database API without authentication.
      tactic: AML.TA0009
      tactic_name: Сбор материалов
      technique: AML.T0036
      technique_name: Данные из информационных репозиториев
    - description: The framework used employee identifiers obtained from the exposed API to test predictable password patterns against the office automation portal. The framework successfully authenticated to the office automation portal using the compromised employee accounts.
      description_line: The framework used employee identifiers obtained from the exposed API to test predictable password patterns against the office automation portal. The framework successfully authenticated to the office automation portal using the compromised employee accounts.
      tactic: AML.TA0004
      tactic_name: Первичный доступ
      technique: AML.T0012
      technique_name: Действующие учетные записи
    - description: The framework systematically tested the 85 compromised office automation accounts against another government information system through an SSO bridge that trusted the existing office automation sessions, providing access to internal dashboards, equipment management interfaces, and personnel statistics pages.
      description_line: The framework systematically tested the 85 compromised office automation accounts against another government information system through an SSO bridge that trusted the existing office automation sessions, providing access to internal dashboards, equipment management interfaces, and personnel statistics pages.
      tactic: AML.TA0015
      tactic_name: Латеральное перемещение
      technique: AML.T0012
      technique_name: Действующие учетные записи
    - description: The framework abused three exposed debug endpoints that accepted arbitrary request bodies and returned authenticated sessions.
      description_line: The framework abused three exposed debug endpoints that accepted arbitrary request bodies and returned authenticated sessions.
      tactic: AML.TA0004
      tactic_name: Первичный доступ
      technique: AML.T0049
      technique_name: Эксплуатация приложения, доступного из интернета
    - description: The framework found a public-facing API that accepted unsigned JWTs with `alg=none`, allowing identity tokens to be forged without the signing key.
      description_line: The framework found a public-facing API that accepted unsigned JWTs with `alg=none`, allowing identity tokens to be forged without the signing key.
      tactic: AML.TA0004
      tactic_name: Первичный доступ
      technique: AML.T0049
      technique_name: Эксплуатация приложения, доступного из интернета
    - description: Using the acquired accesses, the framework automatically retrieved and aggregated the reported personnel, account, configuration, credential, and network information.
      description_line: Using the acquired accesses, the framework automatically retrieved and aggregated the reported personnel, account, configuration, credential, and network information.
      tactic: AML.TA0009
      tactic_name: Сбор материалов
      technique: AML.T0126
      technique_name: Automated Collection
    - description: Over 2,564 personnel records, a complete user database, and internal architecture details were exfiltrated.
      description_line: Over 2,564 personnel records, a complete user database, and internal architecture details were exfiltrated.
      tactic: AML.TA0010
      tactic_name: Эксфильтрация
      technique: AML.T0025
      technique_name: Эксфильтрация киберсредствами
procedure_count: 12
references:
    - title: Inside a Multi-Agent AI Framework Used to Compromise Government Entities in Asia
      url: https://www.dreamgroup.com/blog/inside-a-multi-agent-ai-framework-used-to-compromise-government-entities-in-asia
    - title: Overseas hackers launch AI-agent attacks against government agencies; MODA activates coordinated response and strengthens cyber defenses
      url: https://moda.gov.tw/ACS/press/news/press/20394
    - title: China-linked hackers hit Taiwan in unprecedented 'autonomous' AI cyber attack
      url: https://www.ft.com/content/7d2ab3e0-9085-48f6-b38a-d90260d58795
reporter: Dream Research Labs
source_name: Multi-Agent Framework Compromises Taiwanese Government Systems
target: Taiwanese government agencies and connected government systems
title: Multi-Agent Framework Compromises Taiwanese Government Systems
url: /studies/AML.CS0071/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

In early July 2026, an unknown Chinese-language operator used a multi-agent framework built on Hermes and OpenClaw against government systems that subsequent public reporting identified as Taiwanese.[[financial-times-taiwan-ai-attack]] Dream Research Labs recovered a 160 MB operational workspace containing 1,395 files documenting 12 attack waves conducted from July 1 through July 4.[[dream-multi-agent-framework]] Taiwan's Ministry of Digital Affairs separately confirmed[[moda-ai-agent-attacks]] detecting abnormal attacks during July involving a hybrid of human operation and OpenClaw-assisted activity.

The agentic AI framework coordinated up to eight specialized sub-agents concurrently across reconnaissance, authentication attacks, API testing, vulnerability research, and exploitation. A probabilistic decision engine ranked findings and 14 candidate attack paths, allocated additional testing to promising results, discarded invalidated paths, and used after-action reports to redirect subsequent activity.

Starting from an internet-facing government portal, the framework decompiled client-side application bundles and mapped connected systems, identity infrastructure, and exposed APIs. It obtained access through exposed debug endpoints, unsigned JWT acceptance, and password spraying based on personnel identifiers collected from unauthenticated APIs. Tesseract OCR automated CAPTCHA solving, reportedly helping compromise 85 accounts, 84 of which authenticated to another government system through an SSO bridge without additional MFA or user confirmation.

Dream Research Labs reported the extraction of more than 2,564 personnel records, a complete user-database export, SSO configuration and client information, database credentials, and internal network ranges.

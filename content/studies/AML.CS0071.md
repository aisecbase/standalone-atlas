---
actor: Unknown Chinese-language actor
atlas_id: AML.CS0071
atlas_type: case-study
case_study_type: incident
description: В начале июля 2026 года неизвестный китайскоязычный оператор применил мультиагентный фреймворк на базе Hermes и OpenClaw против государственных систем. В последующих открытых публикациях сообщалось, что речь шла о...
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
      technique_name: Автономная оркестрация атаки
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
      technique_name: Прямая коммуникация агентов
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
      technique_name: Автоматизированный сбор материалов
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
title: Мультиагентный фреймворк скомпрометировал государственные системы Тайваня
url: /studies/AML.CS0071/
---

В начале июля 2026 года неизвестный китайскоязычный оператор применил мультиагентный фреймворк на базе Hermes и OpenClaw против государственных систем. В последующих открытых публикациях сообщалось, что речь шла о системах Тайваня.[[financial-times-taiwan-ai-attack]] Dream Research Labs восстановила рабочее пространство объёмом 160 MB, использовавшееся для проведения операции. Оно содержало 1 395 файлов, в которых были зафиксированы 12 волн атак, проведённых с 1 по 4 июля.[[dream-multi-agent-framework]] Министерство цифрового развития Тайваня отдельно подтвердило[[moda-ai-agent-attacks]], что в июле обнаружило аномальные атаки, сочетавшие ручное управление и действия при поддержке OpenClaw.

Мультиагентный ИИ-фреймворк одновременно координировал работу до восьми специализированных субагентов, занимавшихся разведкой, атаками на механизмы аутентификации, тестированием API, исследованием и эксплуатацией уязвимостей. Вероятностный механизм принятия решений ранжировал выявленные сведения и 14 возможных путей атаки, направлял перспективные результаты на дополнительную проверку, отбрасывал не подтвердившиеся пути и корректировал дальнейшие действия с учётом отчётов по итогам работы.

Начав с государственного портала, доступного из интернета, фреймворк декомпилировал пакеты клиентского кода приложений и составил карту связанных систем, инфраструктуры идентификации и доступных извне API. Он получил доступ благодаря доступным извне отладочным эндпоинтам, приёму неподписанных JWT и распылению паролей с использованием идентификаторов сотрудников, собранных через API без аутентификации. Tesseract OCR автоматизировал решение CAPTCHA, что, согласно отчёту, помогло скомпрометировать 85 учётных записей. С помощью 84 из них удалось пройти аутентификацию в другой государственной системе через SSO-мост без дополнительной MFA или подтверждения со стороны пользователя.

По данным Dream Research Labs, в ходе атаки злоумышленник извлёк более 2 564 записей о сотрудниках, полностью выгрузил базу данных пользователей и получил конфигурацию SSO, сведения о клиентах, учётные данные баз данных и диапазоны внутренних сетей.

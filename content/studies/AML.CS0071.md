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
    - description: В ходе 12 волн атак фреймворк параллельно задействовал до восьми специализированных агентов. Он назначал отдельные задачи по разведке, проверке аутентификации, тестированию API, исследованию уязвимостей, получению учётных данных и эксплуатации уязвимостей; направлял перспективные результаты на дополнительную проверку; запрашивал независимую проверку; объединял отчёты по итогам действий и перенаправлял последующую работу с учётом состояния связанных направлений.
      description_line: В ходе 12 волн атак фреймворк параллельно задействовал до восьми специализированных агентов. Он назначал отдельные задачи по разведке, проверке аутентификации, тестированию API, исследованию уязвимостей, получению учётных данных и эксплуатации уязвимостей; направлял перспективные результаты на дополнительную проверку; запрашивал независимую проверку; объединял отчёты по итогам действий и перенаправлял последующую работу с учётом состояния связанных направлений.
      tactic: AML.TA0001
      tactic_name: Подготовка атаки на ИИ
      technique: AML.T0124
      technique_name: Автономная оркестрация атаки
    - description: Фреймворк сформировал множество возможных многоэтапных путей атаки с учётом подтверждённых необходимых условий, выявленных препятствий и оценок вероятности успеха. Он повышал приоритет путей, подкреплённых проверенными данными, ставил в очередь пути, требовавшие дополнительного исследования, отбрасывал ложноположительные результаты и заблокированные пути, а когда существующие методы не срабатывали — запускал циклы обучения применительно к конкретному целевому объекту.
      description_line: Фреймворк сформировал множество возможных многоэтапных путей атаки с учётом подтверждённых необходимых условий, выявленных препятствий и оценок вероятности успеха. Он повышал приоритет путей, подкреплённых проверенными данными, ставил в очередь пути, требовавшие дополнительного исследования, отбрасывал ложноположительные результаты и заблокированные пути, а когда существующие методы не срабатывали — запускал циклы обучения применительно к конкретному целевому объекту.
      tactic: AML.TA0001
      tactic_name: Подготовка атаки на ИИ
      technique: AML.T0117
      technique_name: Автономная адаптация пути атаки
    - description: Управляющий процесс фреймворка, выполнявший функции оркестратора, обменивался со специализированными субагентами заданиями, обнаруженными сведениями, результатами проверки, сведениями о состоянии и информацией по итогам действий. Сводные результаты учитывались при планировании последующих заданий и волн атак.
      description_line: Управляющий процесс фреймворка, выполнявший функции оркестратора, обменивался со специализированными субагентами заданиями, обнаруженными сведениями, результатами проверки, сведениями о состоянии и информацией по итогам действий. Сводные результаты учитывались при планировании последующих заданий и волн атак.
      tactic: AML.TA0001
      tactic_name: Подготовка атаки на ИИ
      technique: AML.T0118.001
      technique_name: Прямая коммуникация агентов
    - description: 'Фреймворк проводил разведку государственного портала Тайваня, доступного из интернета: анализировал пакеты клиентского кода приложений, прослеживал выявленные связи между элементами инфраструктуры и формировал дополнительные цели разведки. Он выявил 21 связанную систему, шесть подобластей SSO, конфигурацию аутентификации, сведения о ключах подписи и более 36 эндпоинтов API в одной системе.'
      description_line: 'Фреймворк проводил разведку государственного портала Тайваня, доступного из интернета: анализировал пакеты клиентского кода приложений, прослеживал выявленные связи между элементами инфраструктуры и формировал дополнительные цели разведки. Он выявил 21 связанную систему, шесть подобластей SSO, конфигурацию аутентификации, сведения о ключах подписи и более 36 эндпоинтов API в одной системе.'
      tactic: AML.TA0002
      tactic_name: Разведка
      technique: AML.T0116
      technique_name: Автономная разведка
    - description: Фреймворк зондировал основные государственные приложения и API, чтобы выявить доступные извне интерфейсы, особенности аутентификации, ошибки конфигурации и уязвимости. Это сканирование выявило несколько возможных путей проникновения в целевые системы.
      description_line: Фреймворк зондировал основные государственные приложения и API, чтобы выявить доступные извне интерфейсы, особенности аутентификации, ошибки конфигурации и уязвимости. Это сканирование выявило несколько возможных путей проникновения в целевые системы.
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

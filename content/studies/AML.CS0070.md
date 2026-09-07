---
actor: Chinese-speaking threat actor using the aliases knaithe and KnYuan
atlas_id: AML.CS0070
atlas_type: case-study
case_study_type: incident
description: Китайскоязычный злоумышленник, действовавший под псевдонимами knaithe и KnYuan, настроил Hermes Agent на использование DeepSeek в качестве движка рассуждений. Hermes обеспечивал доступ к терминалу, управление со...
generated: true
generated_by: atlasgen
incident_date: "2026-05-07"
incident_date_granularity: Day
incident_date_raw: "2026-05-07"
procedure:
    - description: The actor obtained access to several generative-AI models and services while evaluating an operational toolset. DeepSeek was selected as the primary reasoning engine for the autonomous attack activity.
      description_line: The actor obtained access to several generative-AI models and services while evaluating an operational toolset. DeepSeek was selected as the primary reasoning engine for the autonomous attack activity.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0016.002
      technique_name: Генеративный ИИ
    - description: The actor obtained and configured Hermes Agent as the offensive framework, together with scripts and conventional scanning and exploitation utilities. Hermes provided terminal access, Telegram-based operator control, and a skills system.
      description_line: The actor obtained and configured Hermes Agent as the offensive framework, together with scripts and conventional scanning and exploitation utilities. Hermes provided terminal access, Telegram-based operator control, and a skills system.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0016.001
      technique_name: Программные инструменты
    - description: The actor obtained agent-specific capabilities, including Hermes's framework-bundled godmode skill and the open-source FofaMap MCP server. The MCP server exposed FOFA asset search, natural-language query translation, and Nuclei scan generation to DeepSeek. Unit 42 does not establish that godmode was invoked during the recovered session.
      description_line: The actor obtained agent-specific capabilities, including Hermes's framework-bundled godmode skill and the open-source FofaMap MCP server. The MCP server exposed FOFA asset search, natural-language query translation, and Nuclei scan generation to DeepSeek. Unit 42 does not establish that godmode was invoked during the recovered session.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0016.004
      technique_name: Инструменты ИИ-агента
    - description: The actor created two Hermes skills. web-terminal-exploitation encoded a procedure for unauthenticated WebSocket exploitation, while fofa-cyberspace-search instructed DeepSeek to use the actor's fofoapi.py script for internet asset enumeration. The observed FOFA workflow is consistent with the latter skill; the report does not attribute an action in the recovered session to web-terminal-exploitation.
      description_line: The actor created two Hermes skills. web-terminal-exploitation encoded a procedure for unauthenticated WebSocket exploitation, while fofa-cyberspace-search instructed DeepSeek to use the actor's fofoapi.py script for internet asset enumeration. The observed FOFA workflow is consistent with the latter skill; the report does not attribute an action in the recovered session to web-terminal-exploitation.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0017.002
      technique_name: Инструменты ИИ-агента
    - description: Hermes accessed DeepSeek through its native API and used the model for vulnerability assessment, target selection, command generation, and operational decisions.
      description_line: Hermes accessed DeepSeek through its native API and used the model for vulnerability assessment, target selection, command generation, and operational decisions.
      tactic: AML.TA0000
      tactic_name: Доступ к ИИ-модели
      technique: AML.T0040
      technique_name: Доступ к API инференса ИИ-модели
    - description: After receiving an initial task, DeepSeek sequenced reconnaissance and exploitation actions, evaluated failed prerequisites, abandoned Langflow, compared alternative products and vulnerabilities, and selected n8n. Unit 42 recovered no additional operator input during the session.
      description_line: After receiving an initial task, DeepSeek sequenced reconnaissance and exploitation actions, evaluated failed prerequisites, abandoned Langflow, compared alternative products and vulnerabilities, and selected n8n. Unit 42 recovered no additional operator input during the session.
      tactic: AML.TA0001
      tactic_name: Подготовка атаки на ИИ
      technique: AML.T0117
      technique_name: Автономная адаптация пути атаки
    - description: DeepSeek generated FOFA queries, shell commands, scanner invocations, and direct HTTP probes based on the results returned during the session.
      description_line: DeepSeek generated FOFA queries, shell commands, scanner invocations, and direct HTTP probes based on the results returned during the session.
      tactic: AML.TA0001
      tactic_name: Подготовка атаки на ИИ
      technique: AML.T0102
      technique_name: Генерация вредоносных команд
    - description: DeepSeek investigated Langflow, determined what information and prerequisites were needed, and selected follow-on reconnaissance based on returned results.
      description_line: DeepSeek investigated Langflow, determined what information and prerequisites were needed, and selected follow-on reconnaissance based on returned results.
      tactic: AML.TA0002
      tactic_name: Разведка
      technique: AML.T0116
      technique_name: Автономная разведка
    - description: DeepSeek queried FOFA and obtained records for 84 exposed Langflow instances. These were exposure records, not confirmed vulnerable targets.
      description_line: DeepSeek queried FOFA and obtained records for 84 exposed Langflow instances. These were exposure records, not confirmed vulnerable targets.
      tactic: AML.TA0002
      tactic_name: Разведка
      technique: AML.T0000
      technique_name: Поиск в открытых технических базах данных
    - description: DeepSeek downloaded a public PoC for Langflow CVE-2026-33017. The report does not establish material modification.
      description_line: DeepSeek downloaded a public PoC for Langflow CVE-2026-33017. The report does not establish material modification.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0016.003
      technique_name: Эксплойты
    - description: DeepSeek ran the public Langflow scanner and identified a target running Langflow 1.3.4.
      description_line: DeepSeek ran the public Langflow scanner and identified a target running Langflow 1.3.4.
      tactic: AML.TA0002
      tactic_name: Разведка
      technique: AML.T0006
      technique_name: Активное сканирование
    - description: DeepSeek attempted exploitation, but no target exposed either required prerequisite. No access was obtained.
      description_line: DeepSeek attempted exploitation, but no target exposed either required prerequisite. No access was obtained.
      tactic: AML.TA0004
      tactic_name: Первичный доступ
      technique: AML.T0049
      technique_name: Эксплуатация приложения, доступного из интернета
    - description: DeepSeek assessed Langflow as low value, surveyed exposure across 10 product families, compared vulnerability severity, deployment footprint, PoC availability, and prerequisites, and selected n8n.
      description_line: DeepSeek assessed Langflow as low value, surveyed exposure across 10 product families, compared vulnerability severity, deployment footprint, PoC availability, and prerequisites, and selected n8n.
      tactic: AML.TA0002
      tactic_name: Разведка
      technique: AML.T0116
      technique_name: Автономная разведка
    - description: DeepSeek searched GitHub for trending 2026 CVE PoC repositories sorted by stars.
      description_line: DeepSeek searched GitHub for trending 2026 CVE PoC repositories sorted by stars.
      tactic: AML.TA0002
      tactic_name: Разведка
      technique: AML.T0095.000
      technique_name: Репозитории кода
    - description: DeepSeek downloaded the public n8n PoC chaining CVE-2026-21858 and CVE-2025-68613 and inspected its affected versions and prerequisites.
      description_line: DeepSeek downloaded the public n8n PoC chaining CVE-2026-21858 and CVE-2025-68613 and inspected its affected versions and prerequisites.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0016.003
      technique_name: Эксплойты
    - description: DeepSeek queried FOFA for n8n deployments. FOFA reported 647,017 global results and 25,209 in China; these were not confirmed vulnerable systems.
      description_line: DeepSeek queried FOFA for n8n deployments. FOFA reported 647,017 global results and 25,209 in China; these were not confirmed vulnerable systems.
      tactic: AML.TA0002
      tactic_name: Разведка
      technique: AML.T0000
      technique_name: Поиск в открытых технических базах данных
    - description: DeepSeek sampled approximately 100 Chinese addresses, probed roughly 40 unique systems, identified three running affected versions, inspected form endpoints, and launched parallel scanning against more than 50 remaining targets.
      description_line: DeepSeek sampled approximately 100 Chinese addresses, probed roughly 40 unique systems, identified three running affected versions, inspected form endpoints, and launched parallel scanning against more than 50 remaining targets.
      tactic: AML.TA0002
      tactic_name: Разведка
      technique: AML.T0006
      technique_name: Активное сканирование
    - description: DeepSeek attempted to find and exploit a system meeting the PoC prerequisites. All discovered forms required authentication, and no attempt produced file read, code execution, or initial access.
      description_line: DeepSeek attempted to find and exploit a system meeting the PoC prerequisites. All discovered forms required authentication, and no attempt produced file read, code execution, or initial access.
      tactic: AML.TA0004
      tactic_name: Первичный доступ
      technique: AML.T0049
      technique_name: Эксплуатация приложения, доступного из интернета
procedure_count: 18
references:
    - title: Chinese-Speaking Threat Actor Harnesses AI Models for Autonomous Cyberattacks
      url: https://unit42.paloaltonetworks.com/autonomous-ai-cyber-attack-campaign/
reporter: Palo Alto Networks Unit 42
source_name: Threat Actor Uses a DeepSeek-Powered Hermes Agent in Langflow and n8n Exploitation Attempts
target: Publicly exposed Langflow and n8n systems, primarily in China
title: Злоумышленник использовал Hermes Agent на базе DeepSeek при попытках эксплуатации Langflow и n8n
url: /studies/AML.CS0070/
---

Китайскоязычный злоумышленник, действовавший под псевдонимами knaithe и KnYuan, настроил Hermes Agent на использование DeepSeek в качестве движка рассуждений. Hermes обеспечивал доступ к терминалу, управление со стороны оператора через Telegram и систему навыков, включавшую навыки для работы красной команды — как поставляемые в составе Hermes, так и созданные злоумышленником. Злоумышленник также интегрировал MCP-сервер, предоставлявший возможности поиска через FOFA, перевода запросов и формирования заданий для сканера Nuclei.

Unit 42 восстановила подробную запись сеанса Hermes Agent, датированную 7 мая 2026 года. Злоумышленник передал первоначальное задание через Telegram, однако Unit 42 не обнаружила в записи сеанса никаких дополнительных указаний оператора. После этого Hermes Agent на базе DeepSeek автономно искал цели, получал общедоступные эксплойты, проводил сканирование, оценивал результаты и корректировал свой подход.

Первоначально агент нацелился на Langflow, используя общедоступный эксплойт для CVE-2026-33017. Установив, что доступные цели не соответствовали условиям, необходимым для применения эксплойта, он отказался от этого пути атаки, сравнил уязвимости в 10 семействах продуктов и выбрал n8n с учётом предполагаемой серьёзности уязвимостей, доступности систем извне и возможности их эксплуатации. Затем агент получил общедоступную цепочку эксплойтов для CVE-2026-21858 и CVE-2025-68613 и прозондировал потенциальные целевые системы, выявленные через FOFA. Хотя несколько систем, по всей видимости, работали на затронутых версиях, агент установил, что в них отсутствовала необходимая возможность загружать файлы без аутентификации. Ни одна из автономных попыток эксплуатации не привела к получению доступа.

В отдельных материалах из рабочего пространства были зафиксированы выполнявшиеся вручную действия злоумышленника: извлечение данных из Citrix NetScaler, выполнение команд в Marimo и попытки запуска обратной командной оболочки в системах Tomcat и IKE VPN.

Unit 42 получила возможность наблюдать за этой активностью после того, как Hermes Agent в ответ на команду из Telegram запустил HTTP-сервер для раздачи файлов прямо из домашнего каталога злоумышленника `/home/worker`, а не из изолированного каталога промежуточного хранения. Это непреднамеренно открыло доступ к рабочему пространству злоумышленника, в том числе к конфигурациям ИИ-инструментов, API-ключам, скриптам эксплойтов, спискам целей, истории команд Bash и журналам сеансов автономной эксплуатации.

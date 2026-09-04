---
actor: Chinese-speaking threat actor using the aliases knaithe and KnYuan
atlas_id: AML.CS0070
atlas_type: case-study
case_study_type: incident
description: A Chinese-speaking threat actor operating as knaithe or KnYuan configured Hermes Agent to use DeepSeek as its reasoning engine. Hermes supplied terminal access, Telegram-based operator control, and a skills system...
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
      technique_name: AI Agent Tools
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
title: Threat Actor Uses a DeepSeek-Powered Hermes Agent in Langflow and n8n Exploitation Attempts
url: /studies/AML.CS0070/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

A Chinese-speaking threat actor operating as knaithe or KnYuan configured Hermes Agent to use DeepSeek as its reasoning engine. Hermes supplied terminal access, Telegram-based operator control, and a skills system containing both bundled and actor-created red-team skills. The actor also integrated an MCP server that exposed FOFA search, query translation, and Nuclei scan-generation capabilities.

Unit 42 recovered a detailed Hermes Agent session dated May 7, 2026. The actor supplied an initial task through Telegram, but Unit 42 recovered no additional operator input during the session. The DeepSeek-powered Hermes Agent then autonomously searched for targets, obtained public exploits, ran scans, evaluated the results, and revised its approach.

The agent initially targeted Langflow using a public exploit for CVE-2026-33017. After determining that available targets lacked the exploit's prerequisites, it abandoned that path, compared vulnerabilities across 10 product families, and selected n8n based on apparent severity, exposure, and exploitability. It then obtained a public exploit chain for CVE-2026-21858 and CVE-2025-68613 and probed candidate systems identified through FOFA. Although several systems appeared to run affected versions, the agent found that they lacked the required unauthenticated file-upload functionality. None of the autonomous exploitation attempts obtained access.

Separate workspace evidence documented manual actor activity involving Citrix NetScaler data extraction, Marimo command execution, and reverse-shell attempts against Tomcat and IKE VPN systems.

Unit 42 obtained this visibility after Hermes Agent responded to a Telegram command by starting an HTTP file server from the actor's home directory, `/home/worker`, instead of an isolated staging directory. This unintentionally exposed the actor's workspace, including AI tool configurations, API keys, exploit scripts, target lists, Bash history, and autonomous exploitation session logs.

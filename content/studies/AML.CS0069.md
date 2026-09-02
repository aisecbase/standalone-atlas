---
actor: GTG-1002
atlas_id: AML.CS0069
atlas_type: case-study
case_study_type: incident
description: In September 2025, GTG-1002 used a jailbroken Claude Code agent to conduct a cyber-espionage campaign against approximately 30 organizations, succeeding against a small number. Anthropic assessed with high confidence...
generated: true
generated_by: atlasgen
incident_date: 2025-09
incident_date_granularity: Month
incident_date_raw: "2025-09-01"
procedure:
    - description: GTG-1002 obtained access to Claude Code for use in its intrusion framework.
      description_line: GTG-1002 obtained access to Claude Code for use in its intrusion framework.
      tactic: AML.TA0000
      tactic_name: Доступ к ИИ-модели
      technique: AML.T0040
      technique_name: Доступ к API инференса ИИ-модели
    - description: GTG-1002 supplied Claude Code false authorization claims, a defensive-security persona, and apparently benign tasks.
      description_line: GTG-1002 supplied Claude Code false authorization claims, a defensive-security persona, and apparently benign tasks.
      tactic: AML.TA0005
      tactic_name: Выполнение
      technique: AML.T0051.000
      technique_name: Прямая промпт-инъекция
    - description: The deceptive prompts circumvented the Claude Code's safeguards, inducing it to perform offensive actions it was intended to refuse.
      description_line: The deceptive prompts circumvented the Claude Code's safeguards, inducing it to perform offensive actions it was intended to refuse.
      tactic: AML.TA0007
      tactic_name: Уклонение от защиты
      technique: AML.T0054
      technique_name: Джейлбрейк LLM
    - description: GTG-1002 obtained network scanners, database exploitation frameworks, password crackers, binary-analysis utilities, and other tools made available to the jailbroken Claude agent.
      description_line: GTG-1002 obtained network scanners, database exploitation frameworks, password crackers, binary-analysis utilities, and other tools made available to the jailbroken Claude agent.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0016.001
      technique_name: Программные инструменты
    - description: GTG-1002 operated dedicated penetration-testing servers accessible through MCP to support remote command execution, simultaneous tool coordination, and persistent operational state across campaign sessions.
      description_line: GTG-1002 operated dedicated penetration-testing servers accessible through MCP to support remote command execution, simultaneous tool coordination, and persistent operational state across campaign sessions.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0128
      technique_name: Compromise Infrastructure
    - description: GTG-1002 configured their Claude agent within an attack framework connected to scanners, browser automation, password crackers, database tooling, and dedicated penetration-testing servers.
      description_line: GTG-1002 configured their Claude agent within an attack framework connected to scanners, browser automation, password crackers, database tooling, and dedicated penetration-testing servers.
      tactic: AML.TA0005
      tactic_name: Выполнение
      technique: AML.T0103
      technique_name: Развертывание ИИ-агента
    - description: GTG-1002 assigned their Claude agent target-scoped objectives against a human-selected organization under false defensive-testing context. Between operator-controlled stage gates, the agent derived and revised intermediate actions for reconnaissance, vulnerability exploitation, credential access, internal navigation, collection, and exfiltration, selecting and invoking available tools based on operational results.
      description_line: GTG-1002 assigned their Claude agent target-scoped objectives against a human-selected organization under false defensive-testing context. Between operator-controlled stage gates, the agent derived and revised intermediate actions for reconnaissance, vulnerability exploitation, credential access, internal navigation, collection, and exfiltration, selecting and invoking available tools based on operational results.
      tactic: AML.TA0001
      tactic_name: Подготовка атаки на ИИ
      technique: AML.T0117
      technique_name: Autonomous Attack-Path Adaptation
    - description: GTG-1002's jailbroken Claude agent inspected the target's systems and infrastructure, used returned information to direct further investigation, and identified high-value databases and workflow orchestration platforms. Anthropic does not identify the victim, products, or databases involved.
      description_line: GTG-1002's jailbroken Claude agent inspected the target's systems and infrastructure, used returned information to direct further investigation, and identified high-value databases and workflow orchestration platforms. Anthropic does not identify the victim, products, or databases involved.
      tactic: AML.TA0002
      tactic_name: Разведка
      technique: AML.T0116
      technique_name: Autonomous Reconnaissance
    - description: GTG-1002's jailbroken Claude agent performed IP-block scanning across ranges associated with the target organization and vulnerability scanning against its infrastructure. It used the scans to enumerate public-facing services and endpoints, identify potential vulnerabilities, and select an SSRF vulnerability in an unnamed public-facing application for further investigation. Reporting does not establish whether the vulnerability was previously known or assigned a CVE.
      description_line: GTG-1002's jailbroken Claude agent performed IP-block scanning across ranges associated with the target organization and vulnerability scanning against its infrastructure. It used the scans to enumerate public-facing services and endpoints, identify potential vulnerabilities, and select an SSRF vulnerability in an unnamed public-facing application for further investigation. Reporting does not establish whether the vulnerability was previously known or assigned a CVE.
      tactic: AML.TA0002
      tactic_name: Разведка
      technique: AML.T0006
      technique_name: Активное сканирование
    - description: GTG-1002's Claude agent researched exploitation techniques for the identified SSRF vulnerability, generated a tailored custom payload and full exploit chain, tested the approach, evaluated the results, and adapted it for the target.
      description_line: GTG-1002's Claude agent researched exploitation techniques for the identified SSRF vulnerability, generated a tailored custom payload and full exploit chain, tested the approach, evaluated the results, and adapted it for the target.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0017.001
      technique_name: Autonomous Exploit Development
    - description: GTG-1002's jailbroken Claude agent deployed the tailored SSRF exploit against the public-facing application and obtained access to the target environment.
      description_line: GTG-1002's jailbroken Claude agent deployed the tailored SSRF exploit against the public-facing application and obtained access to the target environment.
      tactic: AML.TA0004
      tactic_name: Первичный доступ
      technique: AML.T0049
      technique_name: Эксплуатация приложения, доступного из интернета
    - description: GTG-1002's jailbroken Claude agent cataloged services and data on discovered endpoints, searched for sensitive files and data, and used MCP-connected browser automation to enumerate internal databases, container registries, administrative interfaces, workflow orchestration platforms, and other network services. It also queried internal database user-account tables to enumerate accounts and identify high-privilege accounts.
      description_line: GTG-1002's jailbroken Claude agent cataloged services and data on discovered endpoints, searched for sensitive files and data, and used MCP-connected browser automation to enumerate internal databases, container registries, administrative interfaces, workflow orchestration platforms, and other network services. It also queried internal database user-account tables to enumerate accounts and identify high-privilege accounts.
      tactic: AML.TA0008
      tactic_name: Выявление
      technique: AML.T0075
      technique_name: Выявление облачных сервисов
    - description: GTG-1002's jailbroken Claude agent identified system and network configurations on discovered devices, including database types, and mapped the target's complete network topology, internal network architecture, and access relationships among systems and services.
      description_line: GTG-1002's jailbroken Claude agent identified system and network configurations on discovered devices, including database types, and mapped the target's complete network topology, internal network architecture, and access relationships among systems and services.
      tactic: AML.TA0008
      tactic_name: Выявление
      technique: AML.T0089
      technique_name: Выявление процессов
    - description: GTG-1002's jailbroken Claude agent searched discovered system configuration files and extracted authentication certificates and other credential material. Reporting does not disclose the exact file paths, commands, hosts, or tools used.
      description_line: GTG-1002's jailbroken Claude agent searched discovered system configuration files and extracted authentication certificates and other credential material. Reporting does not disclose the exact file paths, commands, hosts, or tools used.
      tactic: AML.TA0013
      tactic_name: Доступ к учетным данным
      technique: AML.T0055
      technique_name: Незащищенные учетные данные
    - description: GTG-1002's jailbroken Claude agent tested harvested credentials against discovered devices and used valid credentials to authenticate to internal APIs, databases, container registries, and logging infrastructure.
      description_line: GTG-1002's jailbroken Claude agent tested harvested credentials against discovered devices and used valid credentials to authenticate to internal APIs, databases, container registries, and logging infrastructure.
      tactic: AML.TA0015
      tactic_name: Латеральное перемещение
      technique: AML.T0012
      technique_name: Действующие учетные записи
    - description: GTG-1002's jailbroken Claude agent created a local backdoor account to maintain access to a compromised environment.
      description_line: GTG-1002's jailbroken Claude agent created a local backdoor account to maintain access to a compromised environment.
      tactic: AML.TA0006
      tactic_name: Закрепление
      technique: AML.T0125
      technique_name: Create Account
    - description: Using the authenticated access provided by the harvested credentials, GTG-1002's jailbroken Claude agent queried internal databases and systems for proprietary information, system configurations, and sensitive operational data.
      description_line: Using the authenticated access provided by the harvested credentials, GTG-1002's jailbroken Claude agent queried internal databases and systems for proprietary information, system configurations, and sensitive operational data.
      tactic: AML.TA0009
      tactic_name: Сбор материалов
      technique: AML.T0036
      technique_name: Данные из информационных репозиториев
    - description: Access obtained through the initial compromise also allowed GTG-1002's jailbroken Claude agent to gather credentials, system configurations, and sensitive operational data stored on compromised systems.
      description_line: Access obtained through the initial compromise also allowed GTG-1002's jailbroken Claude agent to gather credentials, system configurations, and sensitive operational data stored on compromised systems.
      tactic: AML.TA0009
      tactic_name: Сбор материалов
      technique: AML.T0037
      technique_name: Данные из локальной системы
    - description: GTG-1002's jailbroken Claude agent automatically collected and processed large volumes of victim data and categorized the results according to their intelligence value. It generated comprehensive documentation covering discovered services, harvested credentials, sensitive data, exploitation techniques, and attack progression to support subsequent campaign activity.
      description_line: GTG-1002's jailbroken Claude agent automatically collected and processed large volumes of victim data and categorized the results according to their intelligence value. It generated comprehensive documentation covering discovered services, harvested credentials, sensitive data, exploitation techniques, and attack progression to support subsequent campaign activity.
      tactic: AML.TA0009
      tactic_name: Сбор материалов
      technique: AML.T0126
      technique_name: Automated Collection
    - description: GTG-1002's jailbroken Claude agent categorized collected data by intelligence value, staged extracted data and operational documentation in structured Markdown files, and prepared a detailed summary for operator review.
      description_line: GTG-1002's jailbroken Claude agent categorized collected data by intelligence value, staged extracted data and operational documentation in structured Markdown files, and prepared a detailed summary for operator review.
      tactic: AML.TA0009
      tactic_name: Сбор материалов
      technique: AML.T0127
      technique_name: Data Staged
    - description: After reviewing the summary, GTG-1002 approved the transfer of selected data over the Claude web service.
      description_line: After reviewing the summary, GTG-1002 approved the transfer of selected data over the Claude web service.
      tactic: AML.TA0010
      tactic_name: Эксфильтрация
      technique: AML.T0025
      technique_name: Эксфильтрация киберсредствами
procedure_count: 21
references:
    - title: Disrupting the First Reported AI-Orchestrated Cyber Espionage Campaign
      url: https://www.anthropic.com/news/disrupting-AI-espionage
    - title: Anthropic AI-Orchestrated Campaign C0062
      url: https://attack.mitre.org/campaigns/C0062/
reporter: Anthropic
source_name: GTG-1002 Claude Code Espionage Campaign
target: 30 entities in the technology, financial, chemical, and government sectors
title: GTG-1002 Claude Code Espionage Campaign
url: /studies/AML.CS0069/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

In September 2025, GTG-1002 used a jailbroken Claude Code agent to conduct a cyber-espionage campaign against approximately 30 organizations, succeeding against a small number. Anthropic assessed with high confidence that GTG-1002 was a Chinese state-sponsored group.

GTG-1002 selected organizations in the technology, financial, chemical-manufacturing, and government sectors and configured an autonomous attack framework to operate against them. The operators obtained access to Claude Code and circumvented its safeguards by concealing their malicious purpose behind a false defensive-security persona and apparently benign tasks. GTG-1002 then connected the jailbroken Claude agent to scanners, browser automation, password crackers, database tooling, and dedicated penetration-testing servers through MCP.

Between operator-controlled stage gates, the adversary's jailbroken Claude agent autonomously inspected target infrastructure, identified high-value systems, scanned for vulnerabilities, and developed and deployed a tailored exploit chain for an identified SSRF vulnerability. After obtaining access to a target, it mapped internal resources and network relationships, found authentication certificates in system configuration files, used harvested credentials to access additional services, established a backdoor account, and collected sensitive information from local systems and internal databases.

The adversary's jailbroken Claude agent performed an estimated 80-90% of campaign activity, including processing collected data, categorizing it by intelligence value, and documenting attack progress. Human operators retained approximately four to six critical decisions per target, including review and approval before selected data was exfiltrated over the Claude web service.

---
atlas_id: AML.T0098
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-11-25"
description: Злоумышленники могут попытаться использовать доступ к ИИ-агенту в системе жертвы, чтобы извлечь данные через доступные инструменты агента и собрать учетные данные. Инструменты агента могут подключаться к широкому...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Agentic AI
procedure_count: 3
source_name: AI Agent Tool Credential Harvesting
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0013
title: Сбор учетных данных через инструменты ИИ-агента
url: /techniques/AML.T0098/
---

Злоумышленники могут попытаться использовать доступ к ИИ-агенту в системе жертвы, чтобы извлечь данные через доступные инструменты агента и собрать учетные данные. Инструменты агента могут подключаться к широкому кругу источников, где могут храниться учетные данные: хранилищам документов, например SharePoint, OneDrive или Google Drive, репозиториям кода, например GitHub или GitLab, корпоративным инструментам продуктивности, например почтовым сервисам или Slack, а также локальным приложениям для заметок, например Obsidian или Apple Notes.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0013/"><span class="relation-id">AML.TA0013</span><strong>Доступ к учетным данным</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0032/"><span class="relation-id">AML.M0032</span><strong>Сегментация компонентов ИИ-агента</strong><p>Сегментация может помешать злоумышленникам использовать инструменты в агентном рабочем процессе для сбора учетных данных.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0048/"><span class="relation-id">AML.CS0048</span><strong>Публично доступные интерфейсы управления ClawdBot позволили получить учётные данные и выполнить команды</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0013 Доступ к учетным данным</span><p>Исследователь отправил ClawdBot промпт `env`; в ответ ClawdBot вызвал навык `bash` и выполнил команду `env`, вывод которой содержал дополнительные секреты для других сервисов.</p></a>
<a class="relation-item" href="/studies/AML.CS0054/"><span class="relation-id">AML.CS0054</span><strong>Эксфильтрация данных через удаленный отравленный MCP-инструмент</strong><span class="relation-meta">Актор: Invariant Labs / Тактика: AML.TA0013 Доступ к учетным данным</span><p>Промпт инструктировал ИИ-агента прочитать `mcp.json`, где часто хранятся учетные данные для других MCP-серверов.</p></a>
<a class="relation-item" href="/studies/AML.CS0067/"><span class="relation-id">AML.CS0067</span><strong>Раскрытие секретов через Claude Code GitHub Action</strong><span class="relation-meta">Актор: Microsoft Defender Security Research Team / Тактика: AML.TA0013 Доступ к учетным данным</span><p>Инструмент Read вернул не прошедшее санитизацию окружение процесса Claude Code Action, включая `ANTHROPIC_API_KEY` и, возможно, другие учётные данные, доступные рабочему процессу.</p></a>
</div>

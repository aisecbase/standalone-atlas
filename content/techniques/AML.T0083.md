---
atlas_id: AML.T0083
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-09-30"
description: Злоумышленники могут получать из конфигурации ИИ-агента учетные данные других инструментов или сервисов в системе. ИИ-агенты часто используют внешние инструменты или сервисы для выполнения действий, например для...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Agentic AI
procedure_count: 2
source_name: Credentials from AI Agent Configuration
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0013
title: Учетные данные из конфигурации ИИ-агента
url: /techniques/AML.T0083/
---

Злоумышленники могут получать из конфигурации ИИ-агента учетные данные других инструментов или сервисов в системе.

ИИ-агенты часто используют внешние инструменты или сервисы для выполнения действий, например для запросов к базам данных, вызова API или взаимодействия с облачными ресурсами. Чтобы обеспечить эти функции, учетные данные, такие как API-ключи, токены и строки подключения, часто хранятся в конфигурационных файлах. Хотя для хранения таких учетных данных и управления ими можно использовать безопасные методы, например выделенные менеджеры секретов или зашифрованные хранилища, на практике их часто размещают в менее защищенных местах ради удобства или простоты развертывания. Если злоумышленник может прочитать или извлечь эти конфигурации, он может получить действительные учетные данные, которые дают прямой доступ к чувствительным системам вне самого агента.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0013/"><span class="relation-id">AML.TA0013</span><strong>Доступ к учетным данным</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0045/"><span class="relation-id">AML.CS0045</span><strong>Эксфильтрация данных через MCP-сервер, используемый Cursor</strong><span class="relation-meta">Актор: Backslash Security Research Team / Тактика: AML.TA0013 Доступ к учетным данным</span><p>Команда оболочки нашла файлы учетных данных `.openapi.apiKey` и `.cursor/mcp.json`, входившие в конфигурацию Cursor.</p></a>
<a class="relation-item" href="/studies/AML.CS0048/"><span class="relation-id">AML.CS0048</span><strong>Публично доступные интерфейсы управления ClawdBot позволили получить учётные данные и выполнить команды</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0013 Доступ к учетным данным</span><p>Исследователь получил доступ к учетным данным разных сервисов, которые хранились в открытом виде в конфигурационном файле ClawdBot `~/.clawdbot/clawdbot.json`; этот файл виден в панели управления ClawdBot. В разных открытых экземплярах ClawdBot он обнаружил ключи API Anthropic, токены Telegram-ботов, учетные данные Slack OAuth и URI привязки устройств Signal.</p></a>
</div>

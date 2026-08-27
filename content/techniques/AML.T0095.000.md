---
atlas_id: AML.T0095.000
atlas_type: technique
attack_ref_id: T1593.003
attack_ref_url: https://attack.mitre.org/techniques/T1593/003/
created_date: "2026-04-22"
description: Злоумышленники могут искать в публичных репозиториях кода информацию о жертве или системе жертвы, которую можно использовать при выборе цели или подготовке атаки. Жертвы могут хранить код или артефакты, связанные с их...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 2
source_name: Code Repositories
subtechnique_count: 0
subtechnique_of: AML.T0095
tactics:
    - AML.TA0002
title: Репозитории кода
url: /techniques/AML.T0095.000/
---

Злоумышленники могут искать в публичных репозиториях кода информацию о жертве или системе жертвы, которую можно использовать при выборе цели или подготовке атаки. Жертвы могут хранить код или артефакты, связанные с их ИИ-системами, в репозиториях на различных сторонних сайтах, таких как GitHub, GitLab, SourceForge и BitBucket. Злоумышленники могут искать в репозиториях кода распространенных ИИ-инструментов, фреймворков, моделей или агентных систем, которые используются жертвой, но не принадлежат ей.

Публичные репозитории кода часто могут быть источником различной информации о жертвах, например об используемых ИИ-фреймворках, библиотеках, моделях, наборах данных, ИИ-агентах и инструментах ИИ-агентов, а также об именах сотрудников. Злоумышленники также могут выявить более чувствительные данные, включая случайно раскрытые учетные данные или API-ключи (например, [учетные данные из конфигурации ИИ-агента](/techniques/AML.T0083)). Информация из этих источников может раскрыть возможности для других форм [разведки](/tactics/AML.TA0002) (например, [сбора целей, индексируемых RAG](/techniques/AML.T0064)), создания операционных ресурсов (например, [приобретения публичных ИИ-артефактов](/techniques/AML.T0002)), [выявления](/tactics/AML.TA0008) (например, [выявления конфигурации ИИ-агента](/techniques/AML.T0084)) и/или [первичного доступа](/tactics/AML.TA0004) (например, [действительных учетных записей](/techniques/AML.T0012) или [фишинга](/techniques/AML.T0052)).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0002/"><span class="relation-id">AML.TA0002</span><strong>Разведка</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0095/"><span class="relation-id">AML.T0095</span><strong>Поиск на открытых сайтах и доменах</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0051/"><span class="relation-id">AML.CS0051</span><strong>Использование OpenClaw для командования и управления через промпт-инъекцию</strong><span class="relation-meta">Актор: HiddenLayer / Тактика: AML.TA0002 Разведка</span><p>Исследователи определили [GitHub-репозиторий OpenClaw](https://github.com/openclaw/openclaw) как источник конфигурационных файлов агента.</p></a>
<a class="relation-item" href="/studies/AML.CS0067/"><span class="relation-id">AML.CS0067</span><strong>Раскрытие секретов через Claude Code GitHub Action</strong><span class="relation-meta">Актор: Microsoft Defender Security Research Team / Тактика: AML.TA0002 Разведка</span><p>The researchers analyzed the Claude Code Action codebase and the obfuscated Claude Agent SDK. They used the implementation details to understand how agent tools executed and where security boundaries were applied.</p></a>
</div>

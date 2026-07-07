---
atlas_id: AML.T0003
atlas_type: technique
attack_ref_id: T1594
attack_ref_url: https://attack.mitre.org/techniques/T1594/
created_date: "2021-05-13"
description: Злоумышленники могут искать на сайтах, принадлежащих жертве, информацию, которую можно использовать при выборе целей. Сайты организации-жертвы могут содержать технические сведения о продуктах или сервисах с поддержкой...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 2
source_name: Search Victim-Owned Websites
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0002
title: Поиск на сайтах организации-жертвы
url: /techniques/AML.T0003/
---

Злоумышленники могут искать на сайтах, принадлежащих жертве, информацию, которую можно использовать при выборе целей.
Сайты организации-жертвы могут содержать технические сведения о продуктах или сервисах с поддержкой ИИ.
На сайтах организации-жертвы могут быть различные сведения, включая названия департаментов или подразделений, физические адреса и данные о ключевых сотрудниках, такие как имена, роли и контактная информация.
Эти сайты также могут содержать сведения о бизнес-операциях и отношениях с партнерами.

Злоумышленники могут искать на сайтах организации-жертвы практически применимую информацию.
Эта информация может помочь злоумышленникам адаптировать атаки, например [состязательные атаки на ИИ](/techniques/AML.T0017.000) или [ручную модификацию](/techniques/AML.T0043.003).
Информация из этих источников может раскрыть возможности для других форм разведки, например [поиска в открытых технических базах данных](/techniques/AML.T0000) или [поиска открытых материалов по анализу уязвимостей ИИ](/techniques/AML.T0001).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0002/"><span class="relation-id">AML.TA0002</span><strong>Разведка</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0000/"><span class="relation-id">AML.M0000</span><strong>Ограничение публичного раскрытия информации</strong><p>Ограничьте публикацию технической информации о продуктах с поддержкой ML и организационной информации о командах, сопровождающих такие продукты.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0014/"><span class="relation-id">AML.CS0014</span><strong>Сбивание с толку антивирусных нейронных сетей</strong><span class="relation-meta">Актор: Kaspersky ML Research Team / Тактика: AML.TA0002 Разведка</span><p>Использование компанией Kaspersky антивирусных детекторов на основе ML публично описано на сайте компании. На практике злоумышленник мог бы использовать эту информацию для выбора цели.</p></a>
<a class="relation-item" href="/studies/AML.CS0039/"><span class="relation-id">AML.CS0039</span><strong>Living Off AI: промпт-инъекция через Jira Service Management</strong><span class="relation-meta">Актор: Cato CTRL / Тактика: AML.TA0002 Разведка</span><p>Исследователи провели разведку, чтобы узнать о сервере Model Context Protocol (MCP) Atlassian и его интеграции с платформой Jira Service Management (JSM). Atlassian предлагает MCP-сервер, который встраивает ИИ в корпоративные рабочие процессы. MCP Atlassian поддерживает ряд действий на основе ИИ, включая суммаризацию обращений, автоматические ответы, классификацию и интеллектуальные рекомендации в JSM и Confluence. Он позволяет инженерам поддержки и внутренним пользователям взаимодействовать с ИИ прямо из привычных интерфейсов.</p></a>
</div>

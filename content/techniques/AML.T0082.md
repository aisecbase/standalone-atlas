---
atlas_id: AML.T0082
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-09-30"
description: Злоумышленники могут пытаться использовать свой доступ к большой языковой модели (LLM) в системе жертвы для сбора учетных данных. Учетные данные могут храниться во внутренних документах, которые могут непреднамеренно...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 1
source_name: RAG Credential Harvesting
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0013
title: Сбор учетных данных через RAG
url: /techniques/AML.T0082/
---

Злоумышленники могут пытаться использовать свой доступ к большой языковой модели (LLM) в системе жертвы для сбора учетных данных. Учетные данные могут храниться во внутренних документах, которые могут непреднамеренно попасть в базу данных RAG, откуда в конечном итоге их может извлечь ИИ-агент.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0013/"><span class="relation-id">AML.TA0013</span><strong>Доступ к учетным данным</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0026/"><span class="relation-id">AML.M0026</span><strong>Настройка разрешений привилегированного ИИ-агента</strong><p>Надлежащий контроль доступа для привилегированных ИИ-агентов может ограничить способность злоумышленника собирать учетные данные из баз данных RAG при компрометации агента.</p></a>
<a class="relation-item" href="/mitigations/AML.M0027/"><span class="relation-id">AML.M0027</span><strong>Настройка разрешений ИИ-агента одного пользователя</strong><p>Настройка ИИ-агентов с разрешениями, унаследованными от пользователя, может ограничить способность злоумышленника собирать учетные данные из баз данных RAG при компрометации агента.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0035/"><span class="relation-id">AML.CS0035</span><strong>Эксфильтрация данных из Slack AI через косвенную промпт-инъекцию</strong><span class="relation-meta">Актор: PromptArmor / Тактика: AML.TA0013 Доступ к учетным данным</span><p>Поскольку Slack AI имеет доступ к приватным каналам пользователя-жертвы, он извлекает API-ключ жертвы.</p></a>
</div>

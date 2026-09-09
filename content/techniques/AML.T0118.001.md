---
atlas_id: AML.T0118.001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-08-31"
description: Автономные ИИ-агенты могут напрямую обмениваться информацией через интерфейсы взаимодействия друг с другом, с субагентами или с оркестратором. Агент может передавать другому агенту контекст операции, обнаруженные...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-08-31"
platforms:
    - Agentic AI
    - Enterprise
procedure_count: 1
source_name: Direct Agent Communication
subtechnique_count: 0
subtechnique_of: AML.T0118
tactics:
    - AML.TA0001
title: Прямая коммуникация агентов
url: /techniques/AML.T0118.001/
---

Автономные ИИ-агенты могут напрямую обмениваться информацией через интерфейсы взаимодействия друг с другом, с субагентами или с оркестратором. Агент может передавать другому агенту контекст операции, обнаруженные сведения, цели, задания, возможности, учётные данные или ограничения, а в ответ получать сведения о состоянии, результаты или выполненную работу.

Прямая коммуникация может включать делегирование, при котором отправляющий агент формулирует или выбирает цель либо подзадачу, а агент-получатель сохраняет существенную свободу действий при выборе способа её выполнения. При прямом обмене агенты также могут сообщать об обнаруженных сведениях, запрашивать независимую проверку, синхронизировать действия, передавать возможности или возвращать результаты, не делегируя при этом новую задачу.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0001/"><span class="relation-id">AML.TA0001</span><strong>Подготовка атаки на ИИ</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0118/"><span class="relation-id">AML.T0118</span><strong>Коммуникация автономных ИИ-агентов</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0118.000/"><span class="relation-id">AML.T0118.000</span><strong>Коммуникация через общие артефакты</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0071/"><span class="relation-id">AML.CS0071</span><strong>Мультиагентный фреймворк скомпрометировал государственные системы Тайваня</strong><span class="relation-meta">Актор: Unknown Chinese-language actor / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Управляющий процесс фреймворка, выполнявший функции оркестратора, обменивался со специализированными субагентами заданиями, обнаруженными сведениями, результатами проверки, сведениями о состоянии и информацией по итогам действий. Сводные результаты учитывались при планировании последующих заданий и волн атак.</p></a>
</div>

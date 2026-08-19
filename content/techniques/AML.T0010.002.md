---
atlas_id: AML.T0010.002
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Данные являются для злоумышленников ключевым вектором компрометации цепочки поставок. Любой ИИ-проект требует данных в той или иной форме. Многие проекты опираются на крупные общедоступные открытые наборы данных....
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 5
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 2
source_name: Data
subtechnique_count: 0
subtechnique_of: AML.T0010
tactics:
    - AML.TA0004
title: Данные
url: /techniques/AML.T0010.002/
---

Данные являются для злоумышленников ключевым вектором компрометации цепочки поставок.

Любой ИИ-проект требует данных в той или иной форме.

Многие проекты опираются на крупные общедоступные открытые наборы данных.

Злоумышленник может рассчитывать на компрометацию этих источников данных.

Вредоносные данные могут быть результатом [отравления обучающих данных](/techniques/AML.T0020) или включать традиционное вредоносное ПО.

Злоумышленник также может нацеливаться на закрытые наборы данных на этапе разметки.

Создание закрытых наборов данных часто требует привлечения внешних сервисов разметки.

Злоумышленник может отравить набор данных, изменяя метки, создаваемые сервисом разметки.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0004/"><span class="relation-id">AML.TA0004</span><strong>Первичный доступ</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0010/"><span class="relation-id">AML.T0010</span><strong>Компрометация цепочки поставок ИИ</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0010.000/"><span class="relation-id">AML.T0010.000</span><strong>Аппаратное обеспечение</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.001/"><span class="relation-id">AML.T0010.001</span><strong>ПО для ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.003/"><span class="relation-id">AML.T0010.003</span><strong>Модель</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.004/"><span class="relation-id">AML.T0010.004</span><strong>Реестр контейнеров</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.005/"><span class="relation-id">AML.T0010.005</span><strong>Инструмент ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0005/"><span class="relation-id">AML.M0005</span><strong>Контроль доступа к ИИ-моделям и хранимым данным</strong><p>Контроль доступа может предотвращать подмену ML-артефактов и несанкционированное копирование.</p></a>
<a class="relation-item" href="/mitigations/AML.M0007/"><span class="relation-id">AML.M0007</span><strong>Санитизация обучающих данных</strong><p>Выявляйте и удаляйте отравленные данные либо устраняйте их последствия, чтобы избежать состязательного дрейфа модели или бэкдор-атак.</p></a>
<a class="relation-item" href="/mitigations/AML.M0014/"><span class="relation-id">AML.M0014</span><strong>Проверка ИИ-артефактов</strong><p>Внедрите надлежащую проверку подписей, чтобы небезопасные ИИ-данные не попадали в систему.</p></a>
<a class="relation-item" href="/mitigations/AML.M0025/"><span class="relation-id">AML.M0025</span><strong>Поддержание происхождения наборов данных ИИ</strong><p>Сведения о происхождении наборов данных могут защищать от компрометации данных в цепочке поставок.</p></a>
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>AI Red Team</strong><p>Introduce controlled untrusted datasets through representative acquisition and ingestion paths. Verify provenance, integrity, sanitization, review, and rejection controls.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0002/"><span class="relation-id">AML.CS0002</span><strong>Отравление VirusTotal</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0004 Первичный доступ</span><p>Злоумышленник загрузил мутированные образцы на платформу.</p></a>
<a class="relation-item" href="/studies/AML.CS0009/"><span class="relation-id">AML.CS0009</span><strong>Отравление Tay</strong><span class="relation-meta">Актор: 4chan Users / Тактика: AML.TA0004 Первичный доступ</span><p>Бот Tay использовал взаимодействия с пользователями Twitter как обучающие данные, чтобы улучшать свои диалоги. Злоумышленники смогли скоординироваться и использовать эту петлю обратной связи, чтобы исказить поведение Tay.</p></a>
</div>

---
atlas_id: AML.T0031
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут снижать производительность целевой модели с помощью состязательных входных данных, чтобы со временем подорвать доверие к системе. Это может привести к тому, что организация-жертва будет тратить...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 4
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 6
source_name: Erode AI Model Integrity
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0011
title: Нарушение целостности ИИ-модели
url: /techniques/AML.T0031/
---

Злоумышленники могут снижать производительность целевой модели с помощью состязательных входных данных, чтобы со временем подорвать доверие к системе.

Это может привести к тому, что организация-жертва будет тратить время и деньги как на попытки исправить систему, так и на ручное выполнение задач, которые система должна была автоматизировать.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0011/"><span class="relation-id">AML.TA0011</span><strong>Воздействие</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0003/"><span class="relation-id">AML.M0003</span><strong>Повышение устойчивости моделей предиктивного ИИ</strong><p>Модели с повышенной устойчивостью менее подвержены атакам на целостность.</p></a>
<a class="relation-item" href="/mitigations/AML.M0006/"><span class="relation-id">AML.M0006</span><strong>Ансамбли моделей предиктивного ИИ</strong><p>Использование нескольких разных моделей повышает устойчивость к атакам.</p></a>
<a class="relation-item" href="/mitigations/AML.M0010/"><span class="relation-id">AML.M0010</span><strong>Восстановление входных данных предиктивного ИИ</strong><p>Предобработка входных данных модели может предотвратить прохождение вредоносных данных через пайплайн машинного обучения.</p></a>
<a class="relation-item" href="/mitigations/AML.M0015/"><span class="relation-id">AML.M0015</span><strong>Обнаружение состязательных входных данных</strong><p>Встраивайте обнаружение состязательных входных данных в пайплайн до того, как входные данные достигнут модели.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0005/"><span class="relation-id">AML.CS0005</span><strong>Атака на сервисы машинного перевода</strong><span class="relation-meta">Актор: Berkeley Artificial Intelligence Research / Тактика: AML.TA0011 Воздействие</span><p>Состязательные атаки могут вызывать ошибки, которые наносят репутационный ущерб компании, предоставляющей сервис перевода, и снижают доверие пользователей к сервисам на базе ИИ.</p></a>
<a class="relation-item" href="/studies/AML.CS0006/"><span class="relation-id">AML.CS0006</span><strong>Ошибочная конфигурация Clearview AI</strong><span class="relation-meta">Актор: Researchers at spiderSilk / Тактика: AML.TA0011 Воздействие</span><p>В результате будущие выпуски приложения могли быть скомпрометированы, что привело бы к ухудшению работы системы распознавания лиц или появлению в ней вредоносных возможностей.</p></a>
<a class="relation-item" href="/studies/AML.CS0009/"><span class="relation-id">AML.CS0009</span><strong>Отравление Tay</strong><span class="relation-meta">Актор: 4chan Users / Тактика: AML.TA0011 Воздействие</span><p>В результате этой скоординированной атаки диалоговые алгоритмы Tay начали учиться генерировать неприемлемые материалы. Усвоение Tay этой оскорбительной лексики привело к тому, что бот начал повторять ее без запроса при взаимодействии с обычными пользователями.</p></a>
<a class="relation-item" href="/studies/AML.CS0019/"><span class="relation-id">AML.CS0019</span><strong>PoisonGPT</strong><span class="relation-meta">Актор: Mithril Security Researchers / Тактика: AML.TA0011 Воздействие</span><p>Из-за ложной информации в выходных данных пользователи могут потерять доверие к приложению.</p></a>
<a class="relation-item" href="/studies/AML.CS0025/"><span class="relation-id">AML.CS0025</span><strong>Отравление крупномасштабных веб-датасетов: атака split-view</strong><span class="relation-meta">Актор: Researchers from Google Deepmind, ETH Zurich, NVIDIA, Robust Intelligence, and Google / Тактика: AML.TA0011 Воздействие</span><p>Модели, обученные на таком датасете, также будут отравлены, что нарушит их целостность. Исследователи показывают, что для успешной атаки достаточно отравить всего 0,01% данных.</p></a>
<a class="relation-item" href="/studies/AML.CS0064/"><span class="relation-id">AML.CS0064</span><strong>Отравленные шаблоны GGUF: атака на цепочку поставок во время инференса</strong><span class="relation-meta">Актор: Pillar Security, Fujitsu Research of Europe / Тактика: AML.TA0011 Воздействие</span><p>Внедрённая инструкция заставляет модель выдавать правдоподобные, но неверные либо сформированные под влиянием злоумышленника ответы. В отсутствие триггера модель продолжает вести себя нормально, из-за чего компрометацию целостности трудно обнаружить.</p></a>
</div>

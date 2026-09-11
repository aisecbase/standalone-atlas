---
atlas_id: AML.T0043.004
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленник может встроить бэкдор-триггер во входные данные, подаваемые на инференс. Такой триггер может быть незаметен или неочевиден для человека. Эта техника применяется вместе с отравлением ИИ-модели и позволяет...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 5
modified_date: "2026-05-27"
platforms:
    - Predictive AI
procedure_count: 1
source_name: Insert Backdoor Trigger
subtechnique_count: 0
subtechnique_of: AML.T0043
tactics:
    - AML.TA0001
title: Добавление бэкдор-триггера
url: /techniques/AML.T0043.004/
---

Злоумышленник может встроить бэкдор-триггер во входные данные, подаваемые на инференс. Такой триггер может быть незаметен или неочевиден для человека. Эта техника применяется вместе с [отравлением ИИ-модели](/techniques/AML.T0018.000) и позволяет злоумышленнику добиться нужного эффекта в целевой модели.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0001/"><span class="relation-id">AML.TA0001</span><strong>Адаптация атак, связанных с ИИ</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0043/"><span class="relation-id">AML.T0043</span><strong>Создание состязательных данных</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0043.000/"><span class="relation-id">AML.T0043.000</span><strong>Оптимизация в режиме белого ящика</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.001/"><span class="relation-id">AML.T0043.001</span><strong>Оптимизация в режиме чёрного ящика</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.002/"><span class="relation-id">AML.T0043.002</span><strong>Перенос на модель чёрного ящика</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.003/"><span class="relation-id">AML.T0043.003</span><strong>Ручная модификация</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0003/"><span class="relation-id">AML.M0003</span><strong>Повышение устойчивости моделей предиктивного ИИ</strong><p>Модели с повышенной устойчивостью лучше противостоят состязательным входным данным.</p></a>
<a class="relation-item" href="/mitigations/AML.M0006/"><span class="relation-id">AML.M0006</span><strong>Ансамбли моделей предиктивного ИИ</strong><p>Использование ансамбля моделей усложняет создание эффективных состязательных данных и повышает общую устойчивость.</p></a>
<a class="relation-item" href="/mitigations/AML.M0008/"><span class="relation-id">AML.M0008</span><strong>Валидация ИИ-модели</strong><p>Проверка того, что ИИ-модель не реагирует на бэкдор-триггеры, может повысить уверенность в том, что модель не была отравлена.</p></a>
<a class="relation-item" href="/mitigations/AML.M0010/"><span class="relation-id">AML.M0010</span><strong>Восстановление входных данных предиктивного ИИ</strong><p>Восстановление входных данных может помогать исправлять состязательные входные данные.</p></a>
<a class="relation-item" href="/mitigations/AML.M0015/"><span class="relation-id">AML.M0015</span><strong>Обнаружение состязательных входных данных для предиктивного ИИ</strong><p>Встраивайте обнаружение состязательных входных данных, чтобы блокировать вредоносные входные данные во время инференса.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0013/"><span class="relation-id">AML.CS0013</span><strong>Бэкдор-атака на модели глубокого обучения в мобильных приложениях</strong><span class="relation-meta">Актор: Yuanchun Li, Jiayi Hua, Haoyu Wang, Chunyang Chen, Yunxin Liu / Тактика: AML.TA0001 Адаптация атак, связанных с ИИ</span><p>Триггер размещается в физической среде, где его захватывает камера устройства жертвы, после чего он обрабатывается ML-моделью с бэкдором.</p></a>
</div>

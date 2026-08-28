---
atlas_id: AML.T0043.001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: При атаках в режиме чёрного ящика злоумышленник имеет доступ к целевой модели в режиме чёрного ящика, то есть через доступ к API инференса ИИ-модели. При таких атаках злоумышленник может использовать API, который...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 7
modified_date: "2026-05-27"
platforms:
    - Predictive AI
procedure_count: 2
source_name: Black-Box Optimization
subtechnique_count: 0
subtechnique_of: AML.T0043
tactics:
    - AML.TA0001
title: Оптимизация в режиме чёрного ящика
url: /techniques/AML.T0043.001/
---

При атаках в режиме чёрного ящика злоумышленник имеет доступ к целевой модели в режиме чёрного ящика, то есть через [доступ к API инференса ИИ-модели](/techniques/AML.T0040). При таких атаках злоумышленник может использовать API, который находится под мониторингом организации-жертвы. Эти атаки обычно менее эффективны и требуют большего числа инференсов, чем атаки с [оптимизацией в режиме белого ящика](/techniques/AML.T0043.000), но требуют значительно меньшего уровня доступа.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0001/"><span class="relation-id">AML.TA0001</span><strong>Подготовка атаки на ИИ</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0043/"><span class="relation-id">AML.T0043</span><strong>Создание состязательных данных</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0043.000/"><span class="relation-id">AML.T0043.000</span><strong>Оптимизация в режиме белого ящика</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.002/"><span class="relation-id">AML.T0043.002</span><strong>Перенос на модель чёрного ящика</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.003/"><span class="relation-id">AML.T0043.003</span><strong>Ручная модификация</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.004/"><span class="relation-id">AML.T0043.004</span><strong>Добавление бэкдор-триггера</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0002/"><span class="relation-id">AML.M0002</span><strong>Обфускация выходных данных предиктивного ИИ</strong><p>Обфускация выходных данных модели снижает способность злоумышленника создавать эффективные состязательные входные данные.</p></a>
<a class="relation-item" href="/mitigations/AML.M0003/"><span class="relation-id">AML.M0003</span><strong>Усиление устойчивости модели</strong><p>Усиленные модели более устойчивы к состязательным входным данным.</p></a>
<a class="relation-item" href="/mitigations/AML.M0004/"><span class="relation-id">AML.M0004</span><strong>Ограничение количества запросов к ИИ-модели</strong><p>Limit volume of model queries to prevent or slow an adversary&#39;s ability to perform black-box optimization attacks.</p></a>
<a class="relation-item" href="/mitigations/AML.M0006/"><span class="relation-id">AML.M0006</span><strong>Использование ансамблевых методов</strong><p>Использование ансамбля моделей усложняет создание эффективных состязательных данных и повышает общую устойчивость.</p></a>
<a class="relation-item" href="/mitigations/AML.M0010/"><span class="relation-id">AML.M0010</span><strong>Восстановление входных данных</strong><p>Восстановление входных данных добавляет дополнительный слой неопределенности и случайности, когда злоумышленник оценивает связь между входом и выходом.</p></a>
<a class="relation-item" href="/mitigations/AML.M0015/"><span class="relation-id">AML.M0015</span><strong>Обнаружение состязательных входных данных</strong><p>Отслеживайте запросы и паттерны запросов к целевой модели и блокируйте доступ при обнаружении подозрительных запросов.</p></a>
<a class="relation-item" href="/mitigations/AML.M0019/"><span class="relation-id">AML.M0019</span><strong>Контроль доступа к ИИ-моделям и данным в продакшене</strong><p>Контроль доступа к API модели может лишить злоумышленников доступа, необходимого для методов оптимизации в режиме чёрного ящика.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0001/"><span class="relation-id">AML.CS0001</span><strong>Обход обнаружения DGA-доменов ботнетов</strong><span class="relation-meta">Актор: Palo Alto Networks AI Research Team / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Исследователи использовали технику мутации для генерации доменных имен, обходящих обнаружение.</p></a>
<a class="relation-item" href="/studies/AML.CS0011/"><span class="relation-id">AML.CS0011</span><strong>Обход ИИ на периферии Microsoft</strong><span class="relation-meta">Актор: Azure Red Team / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Красная команда создала автоматизированную систему, которая непрерывно изменяла исходное целевое изображение так, чтобы обмануть ML-модель и заставить ее выдавать неверные результаты инференса, при этом возмущения на изображении оставались незаметными для человеческого глаза.</p></a>
</div>

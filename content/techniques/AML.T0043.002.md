---
atlas_id: AML.T0043.002
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: При атаках с переносом на модель чёрного ящика злоумышленник использует одну или несколько прокси-моделей, к которым у него есть полный доступ и которые близки по поведению к целевой модели. Такие прокси-модели могут...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 4
modified_date: "2026-05-27"
platforms:
    - Predictive AI
procedure_count: 3
source_name: Black-Box Transfer
subtechnique_count: 0
subtechnique_of: AML.T0043
tactics:
    - AML.TA0001
title: Перенос на модель чёрного ящика
url: /techniques/AML.T0043.002/
---

При атаках с переносом на модель чёрного ящика злоумышленник использует одну или несколько прокси-моделей, к которым у него есть полный доступ и которые близки по поведению к целевой модели. Такие прокси-модели могут быть обучены через [создание прокси-модели ИИ](/techniques/AML.T0005) или [обучение прокси через репликацию](/techniques/AML.T0005.001). Злоумышленник применяет [оптимизацию в режиме белого ящика](/techniques/AML.T0043.000) к прокси-моделям, чтобы сгенерировать состязательные примеры. Если набор прокси-моделей достаточно близок к целевой модели, состязательный пример должен переноситься с одной модели на другую. Это означает, что атака, работающая против прокси-моделей, с высокой вероятностью сработает и против целевой модели. Если у злоумышленника есть [доступ к API инференса ИИ-модели](/techniques/AML.T0040), он может использовать [проверку атаки](/techniques/AML.T0042), чтобы подтвердить работоспособность атаки и учесть эту информацию в процессе обучения.


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
<a class="relation-item" href="/techniques/AML.T0043.001/"><span class="relation-id">AML.T0043.001</span><strong>Оптимизация в режиме чёрного ящика</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.003/"><span class="relation-id">AML.T0043.003</span><strong>Ручная модификация</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.004/"><span class="relation-id">AML.T0043.004</span><strong>Добавление бэкдор-триггера</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0003/"><span class="relation-id">AML.M0003</span><strong>Повышение устойчивости моделей предиктивного ИИ</strong><p>Модели с повышенной устойчивостью лучше противостоят состязательным входным данным.</p></a>
<a class="relation-item" href="/mitigations/AML.M0006/"><span class="relation-id">AML.M0006</span><strong>Использование ансамблевых методов</strong><p>Использование ансамбля моделей усложняет создание эффективных состязательных данных и повышает общую устойчивость.</p></a>
<a class="relation-item" href="/mitigations/AML.M0010/"><span class="relation-id">AML.M0010</span><strong>Восстановление входных данных</strong><p>Восстановление входных данных может помогать исправлять состязательные входные данные.</p></a>
<a class="relation-item" href="/mitigations/AML.M0015/"><span class="relation-id">AML.M0015</span><strong>Обнаружение состязательных входных данных</strong><p>Встраивайте обнаружение состязательных входных данных, чтобы блокировать вредоносные входные данные во время инференса.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0005/"><span class="relation-id">AML.CS0005</span><strong>Атака на сервисы машинного перевода</strong><span class="relation-meta">Актор: Berkeley Artificial Intelligence Research / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Реплицированные модели использовались для генерации состязательных примеров, которые успешно срабатывали на сервисах машинного перевода с закрытой внутренней логикой.</p></a>
<a class="relation-item" href="/studies/AML.CS0008/"><span class="relation-id">AML.CS0008</span><strong>Обход ProofPoint</strong><span class="relation-meta">Актор: Researchers at Silent Break Security / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Затем исследователи ML алгоритмически нашли в этой &#34;офлайн&#34; прокси-модели образцы, которые помогли получить нужное представление о ее поведении и влиятельных переменных. Примеры образцов с хорошими оценками: &#34;calculation&#34;, &#34;asset&#34; и &#34;tyson&#34;. Примеры образцов с плохими оценками: &#34;software&#34;, &#34;99&#34; и &#34;unsub&#34;.</p></a>
<a class="relation-item" href="/studies/AML.CS0014/"><span class="relation-id">AML.CS0014</span><strong>Сбивание с толку антивирусных нейронных сетей</strong><span class="relation-meta">Актор: Kaspersky ML Research Team / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>С помощью разработанного градиентного алгоритма из вредоносных файлов были созданы состязательные вредоносные файлы для прокси-модели, предназначенные для переноса на целевую модель в режиме чёрного ящика.</p></a>
</div>

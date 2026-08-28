---
atlas_id: AML.T0043.000
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: При оптимизации в режиме белого ящика злоумышленник имеет полный доступ к целевой модели и напрямую оптимизирует состязательный пример. Состязательные примеры, обученные таким способом, наиболее эффективны против...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 6
modified_date: "2026-05-27"
platforms:
    - Predictive AI
procedure_count: 3
source_name: White-Box Optimization
subtechnique_count: 0
subtechnique_of: AML.T0043
tactics:
    - AML.TA0001
title: Оптимизация в режиме белого ящика
url: /techniques/AML.T0043.000/
---

При оптимизации в режиме белого ящика злоумышленник имеет полный доступ к целевой модели и напрямую оптимизирует состязательный пример. Состязательные примеры, обученные таким способом, наиболее эффективны против целевой модели.


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
<a class="relation-item" href="/techniques/AML.T0043.001/"><span class="relation-id">AML.T0043.001</span><strong>Оптимизация в режиме чёрного ящика</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.002/"><span class="relation-id">AML.T0043.002</span><strong>Перенос на модель чёрного ящика</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.003/"><span class="relation-id">AML.T0043.003</span><strong>Ручная модификация</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.004/"><span class="relation-id">AML.T0043.004</span><strong>Добавление бэкдор-триггера</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0003/"><span class="relation-id">AML.M0003</span><strong>Повышение устойчивости моделей предиктивного ИИ</strong><p>Модели с повышенной устойчивостью лучше противостоят состязательным входным данным.</p></a>
<a class="relation-item" href="/mitigations/AML.M0005/"><span class="relation-id">AML.M0005</span><strong>Контроль доступа к ИИ-моделям и хранимым данным</strong><p>Контроль доступа может сократить избыточный доступ к ИИ-моделям и не дать злоумышленнику получить доступ в режиме белого ящика.</p></a>
<a class="relation-item" href="/mitigations/AML.M0006/"><span class="relation-id">AML.M0006</span><strong>Ансамбли моделей предиктивного ИИ</strong><p>Использование ансамбля моделей усложняет создание эффективных состязательных данных и повышает общую устойчивость.</p></a>
<a class="relation-item" href="/mitigations/AML.M0010/"><span class="relation-id">AML.M0010</span><strong>Восстановление входных данных</strong><p>Восстановление входных данных может помогать исправлять состязательные входные данные.</p></a>
<a class="relation-item" href="/mitigations/AML.M0015/"><span class="relation-id">AML.M0015</span><strong>Обнаружение состязательных входных данных</strong><p>Встраивайте обнаружение состязательных входных данных, чтобы блокировать вредоносные входные данные во время инференса.</p></a>
<a class="relation-item" href="/mitigations/AML.M0017/"><span class="relation-id">AML.M0017</span><strong>Методы распространения ИИ-моделей</strong><p>При полном доступе к модели злоумышленник может проводить атаки в режиме белого ящика.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0010/"><span class="relation-id">AML.CS0010</span><strong>Нарушение работы сервиса Microsoft Azure</strong><span class="relation-meta">Актор: Microsoft AI Red Team / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Используя целевую модель и данные, красная команда в офлайн-режиме подготовила состязательные данные, рассчитанные на обход модели.</p></a>
<a class="relation-item" href="/studies/AML.CS0012/"><span class="relation-id">AML.CS0012</span><strong>Обход системы идентификации лиц с помощью физических контрмер</strong><span class="relation-meta">Актор: MITRE AI Red Team / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Используя прокси-модель, красная команда оптимизировала состязательные визуальные паттерны для атаки с физической наклейкой на основе метода expectation over transformation.</p></a>
<a class="relation-item" href="/studies/AML.CS0058/"><span class="relation-id">AML.CS0058</span><strong>Извлечение ИИ-моделей из Google Photos</strong><span class="relation-meta">Актор: Skyld / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Восстановленные модели TensorFlow Lite могли использоваться для генерации состязательных примеров в режиме белого ящика.</p></a>
</div>

---
atlas_id: AML.T0043.003
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут вручную изменять входные данные, чтобы создавать состязательные данные. Они могут использовать свои знания о целевой модели и изменять те части данных, которые, по их предположению, помогают...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 5
modified_date: "2026-05-27"
platforms:
    - Predictive AI
procedure_count: 3
source_name: Manual Modification
subtechnique_count: 0
subtechnique_of: AML.T0043
tactics:
    - AML.TA0001
title: Ручная модификация
url: /techniques/AML.T0043.003/
---

Злоумышленники могут вручную изменять входные данные, чтобы создавать состязательные данные. Они могут использовать свои знания о целевой модели и изменять те части данных, которые, по их предположению, помогают модели выполнять её задачу. Злоумышленник может действовать методом проб и ошибок, пока не сможет подтвердить, что получил работоспособный состязательный вход.


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
<a class="relation-item" href="/techniques/AML.T0043.002/"><span class="relation-id">AML.T0043.002</span><strong>Перенос на модель чёрного ящика</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.004/"><span class="relation-id">AML.T0043.004</span><strong>Добавление бэкдор-триггера</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0003/"><span class="relation-id">AML.M0003</span><strong>Повышение устойчивости моделей предиктивного ИИ</strong><p>Модели с повышенной устойчивостью лучше противостоят состязательным входным данным.</p></a>
<a class="relation-item" href="/mitigations/AML.M0004/"><span class="relation-id">AML.M0004</span><strong>Ограничение объёма и частоты запросов к ИИ-сервису</strong><p>Ограничивайте объём запросов к модели, чтобы помешать злоумышленнику дорабатывать подготовленные вручную состязательные входные данные или замедлить этот процесс.</p></a>
<a class="relation-item" href="/mitigations/AML.M0006/"><span class="relation-id">AML.M0006</span><strong>Использование ансамблевых методов</strong><p>Использование ансамбля моделей усложняет создание эффективных состязательных данных и повышает общую устойчивость.</p></a>
<a class="relation-item" href="/mitigations/AML.M0010/"><span class="relation-id">AML.M0010</span><strong>Восстановление входных данных</strong><p>Восстановление входных данных может помогать исправлять состязательные входные данные.</p></a>
<a class="relation-item" href="/mitigations/AML.M0015/"><span class="relation-id">AML.M0015</span><strong>Обнаружение состязательных входных данных</strong><p>Встраивайте обнаружение состязательных входных данных, чтобы блокировать вредоносные входные данные во время инференса.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0000/"><span class="relation-id">AML.CS0000</span><strong>Обход детектора C&amp;C-трафика вредоносного ПО на основе глубокого обучения</strong><span class="relation-meta">Актор: Palo Alto Networks AI Research Team / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Мы создали образцы для обхода, удалив из заголовка пакета поля, которые обычно не используются для C&amp;C-коммуникации, например `cache-control`, `connection` и т. д.</p></a>
<a class="relation-item" href="/studies/AML.CS0003/"><span class="relation-id">AML.CS0003</span><strong>Обход ИИ-детектора вредоносного ПО Cylance</strong><span class="relation-meta">Актор: Skylight Cyber / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Используя эти знания, исследователи объединили атрибуты заведомо легитимных файлов с вредоносным ПО, чтобы вручную создать состязательные образцы вредоносного ПО.</p></a>
<a class="relation-item" href="/studies/AML.CS0032/"><span class="relation-id">AML.CS0032</span><strong>Попытка обхода ML-системы обнаружения фишинговых веб-страниц</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Наблюдалось несколько простых, но эффективных стратегий ручного изменения логотипов: Стиль названия компании - 25; Размытый логотип - 23; Обрезка - 20; Без названия компании - 16; Без визуального логотипа - 13; Другой визуальный логотип - 12; Растягивание логотипа - 11; Несколько форм — изображения - 10; Фоновые узоры - 8; Обфускация страницы входа - 6; Маскирование - 3</p></a>
</div>

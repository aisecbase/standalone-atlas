---
atlas_id: AML.T0002.000
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут собирать публичные наборы данных для использования в своих операциях. Наборы данных, используемые организацией-жертвой, или наборы данных, репрезентативные по отношению к данным...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 6
source_name: Datasets
subtechnique_count: 0
subtechnique_of: AML.T0002
tactics:
    - AML.TA0003
title: Наборы данных
url: /techniques/AML.T0002.000/
---

Злоумышленники могут собирать публичные наборы данных для использования в своих операциях.
Наборы данных, используемые организацией-жертвой, или наборы данных, репрезентативные по отношению к данным организации-жертвы, могут быть ценны для злоумышленников.
Наборы данных могут храниться в облачных хранилищах или на сайтах, принадлежащих жертве.
Для доступа к некоторым наборам данных злоумышленнику необходимо [создать учетные записи](/techniques/AML.T0021).

Полученные наборы данных помогают злоумышленнику развивать операции, подготавливать атаки и адаптировать атаки под организацию-жертву.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0002/"><span class="relation-id">AML.T0002</span><strong>Получение публичных ИИ-артефактов</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0002.001/"><span class="relation-id">AML.T0002.001</span><strong>Модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0002.002/"><span class="relation-id">AML.T0002.002</span><strong>Конфигурация ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0001/"><span class="relation-id">AML.M0001</span><strong>Ограничение публикации артефактов модели</strong><p>Ограничение публикации наборов данных может снизить способность злоумышленника нацеливаться на продакшен-модели, обученные на тех же или похожих данных.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0000/"><span class="relation-id">AML.CS0000</span><strong>Обход детектора C&amp;C-трафика вредоносного ПО на основе глубокого обучения</strong><span class="relation-meta">Актор: Palo Alto Networks AI Research Team / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Мы получили набор данных HTTP-трафика командования и управления, состоящий примерно из 33 млн безвредных и 27 млн вредоносных заголовков HTTP-пакетов.</p></a>
<a class="relation-item" href="/studies/AML.CS0005/"><span class="relation-id">AML.CS0005</span><strong>Атака на сервисы машинного перевода</strong><span class="relation-meta">Актор: Berkeley Artificial Intelligence Research / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи собрали похожие наборы данных, которые использовали целевые сервисы машинного перевода.</p></a>
<a class="relation-item" href="/studies/AML.CS0007/"><span class="relation-id">AML.CS0007</span><strong>Репликация модели GPT-2</strong><span class="relation-meta">Актор: Researchers at Brown University / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Используя собранную документацию, исследователи смогли вручную воссоздать набор данных, применявшийся в исходной статье о GPT-2.</p></a>
<a class="relation-item" href="/studies/AML.CS0012/"><span class="relation-id">AML.CS0012</span><strong>Обход системы идентификации лиц с помощью физических контрмер</strong><span class="relation-meta">Актор: MITRE AI Red Team / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Команда получила репрезентативные данные из открытых источников.</p></a>
<a class="relation-item" href="/studies/AML.CS0014/"><span class="relation-id">AML.CS0014</span><strong>Сбивание с толку антивирусных нейронных сетей</strong><span class="relation-meta">Актор: Kaspersky ML Research Team / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи собрали набор данных из вредоносных и чистых файлов. Они просканировали этот набор целевым антивирусным решением на основе ML и разметили образцы в соответствии с предсказаниями ML-детектора.</p></a>
<a class="relation-item" href="/studies/AML.CS0025/"><span class="relation-id">AML.CS0025</span><strong>Отравление крупномасштабных веб-датасетов: атака split-view</strong><span class="relation-meta">Актор: Researchers from Google Deepmind, ETH Zurich, NVIDIA, Robust Intelligence, and Google / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи скачивают крупномасштабный веб-датасет, представляющий собой список URL на отдельные элементы данных.</p></a>
</div>

---
atlas_id: AML.T0017.000
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2023-10-25"
description: Злоумышленники могут разрабатывать собственные состязательные атаки. Они могут использовать существующие библиотеки как отправную точку (готовые реализации состязательных атак на ИИ). Они могут реализовывать идеи,...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 5
source_name: Adversarial AI Attacks
subtechnique_count: 0
subtechnique_of: AML.T0017
tactics:
    - AML.TA0003
title: Состязательные атаки на ИИ
url: /techniques/AML.T0017.000/
---

Злоумышленники могут разрабатывать собственные состязательные атаки.

Они могут использовать существующие библиотеки как отправную точку ([готовые реализации состязательных атак на ИИ](/techniques/AML.T0016.000)).

Они могут реализовывать идеи, описанные в публичных научных публикациях, или разрабатывать специализированные атаки против модели организации-жертвы.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0017/"><span class="relation-id">AML.T0017</span><strong>Разработка средств для атаки</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0017.001/"><span class="relation-id">AML.T0017.001</span><strong>Autonomous Exploit Development</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0017.002/"><span class="relation-id">AML.T0017.002</span><strong>AI Agent Tools</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0001/"><span class="relation-id">AML.CS0001</span><strong>Обход обнаружения DGA-доменов ботнетов</strong><span class="relation-meta">Актор: Palo Alto Networks AI Research Team / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи разработали универсальную технику мутации, требующую минимального числа итераций.</p></a>
<a class="relation-item" href="/studies/AML.CS0003/"><span class="relation-id">AML.CS0003</span><strong>Обход ИИ-детектора вредоносного ПО Cylance</strong><span class="relation-meta">Актор: Skylight Cyber / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи использовали информацию о репутационном скоринге, чтобы с помощью обратной разработки определить, какие атрибуты давали тот или иной уровень положительной или отрицательной репутации. В процессе они обнаружили вторичную модель, которая переопределяла первую модель. Положительные оценки вторичной модели переопределяли решение основной ML-модели.</p></a>
<a class="relation-item" href="/studies/AML.CS0013/"><span class="relation-id">AML.CS0013</span><strong>Бэкдор-атака на модели глубокого обучения в мобильных приложениях</strong><span class="relation-meta">Актор: Yuanchun Li, Jiayi Hua, Haoyu Wang, Chunyang Chen, Yunxin Liu / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи разработали новый подход к внедрению бэкдора в скомпилированную модель, который может активироваться визуальным триггером. Они внедряют в модель &#34;нейронную полезную нагрузку&#34;, состоящую из сети обнаружения триггера и условной логики. Детектор триггера обучается распознавать визуальный триггер, который будет размещен в реальном мире. Условная логика позволяет исследователям обходить модель жертвы при обнаружении триггера и выдавать выбранные ими выходные данные модели. Для обучения детектора триггера нужны только общий набор данных той же модальности, что и целевая модель, например ImageNet для классификации изображений, и несколько фотографий нужного триггера.</p></a>
<a class="relation-item" href="/studies/AML.CS0014/"><span class="relation-id">AML.CS0014</span><strong>Сбивание с толку антивирусных нейронных сетей</strong><span class="relation-meta">Актор: Kaspersky ML Research Team / Тактика: AML.TA0003 Подготовка ресурсов</span><p>С помощью реверс-инжиниринга локального экстрактора признаков исследователи смогли собрать сведения о входных признаках, используемых облачным ML-детектором. Модель собирает признаки PE-заголовка, признаки секций и статистику данных секций, а также информацию о строках файла. Был разработан градиентный состязательный алгоритм для исполняемых файлов. Алгоритм изменяет признаки файла, чтобы избежать обнаружения прокси-моделью, сохраняя при этом ту же вредоносную полезную нагрузку.</p></a>
<a class="relation-item" href="/studies/AML.CS0064/"><span class="relation-id">AML.CS0064</span><strong>Отравленные шаблоны GGUF: атака на цепочку поставок во время инференса</strong><span class="relation-meta">Актор: Pillar Security, Fujitsu Research of Europe / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленник разрабатывает бэкдор на основе шаблона, содержащий условную логику срабатывания триггера и полезную нагрузку в виде подконтрольной злоумышленнику инструкции. Триггер выбирается так, чтобы срабатывать при обычном использовании целевого приложения.</p></a>
</div>

---
atlas_id: AML.T0005.001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут реплицировать непубличную модель. Многократно используя доступ к API инференса ИИ-модели организации-жертвы, злоумышленник может собрать результаты инференса целевой модели в набор данных....
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 3
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 2
source_name: Train Proxy via Replication
subtechnique_count: 0
subtechnique_of: AML.T0005
tactics:
    - AML.TA0001
title: Обучение прокси-модели через репликацию
url: /techniques/AML.T0005.001/
---

Злоумышленники могут реплицировать непубличную модель.

Многократно используя [доступ к API инференса ИИ-модели](/techniques/AML.T0040) организации-жертвы, злоумышленник может собрать результаты инференса целевой модели в набор данных.

Полученные результаты инференса используются в качестве меток для офлайн-обучения отдельной модели, которая будет имитировать поведение и показатели качества целевой модели.

Реплицированная модель, с высокой точностью воспроизводящая поведение целевой модели, представляет собой ценный ресурс при подготовке атаки.

Злоумышленник может использовать реплицированную модель для [создания состязательных данных](/techniques/AML.T0043) в различных целях, например для [обхода ИИ-модели](/techniques/AML.T0015) или [зашумления ИИ-системы нерелевантными данными](/techniques/AML.T0046).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0001/"><span class="relation-id">AML.TA0001</span><strong>Адаптация атак, связанных с ИИ</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0005/"><span class="relation-id">AML.T0005</span><strong>Создание прокси-модели ИИ</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0005.000/"><span class="relation-id">AML.T0005.000</span><strong>Обучение прокси-модели на собранных ИИ-артефактах</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0005.002/"><span class="relation-id">AML.T0005.002</span><strong>Использование предварительно обученной модели</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0002/"><span class="relation-id">AML.M0002</span><strong>Обфускация выходных данных предиктивного ИИ</strong><p>Обфускация выходных данных модели ограничивает способность злоумышленника создать точную прокси-модель, запрашивая модель и наблюдая ее выходные данные.</p></a>
<a class="relation-item" href="/mitigations/AML.M0004/"><span class="relation-id">AML.M0004</span><strong>Ограничение объёма и частоты запросов к ИИ-сервису</strong><p>Ограничивайте запросы на инференс, чтобы сократить объём размеченных выходных данных, доступных для обучения прокси-модели.</p></a>
<a class="relation-item" href="/mitigations/AML.M0024/"><span class="relation-id">AML.M0024</span><strong>Логирование телеметрии ИИ</strong><p>Логирование телеметрии может помочь выявить эксфильтрацию обучающего набора данных для прокси-модели.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0005/"><span class="relation-id">AML.CS0005</span><strong>Атака на сервисы машинного перевода</strong><span class="relation-meta">Актор: Berkeley Artificial Intelligence Research / Тактика: AML.TA0001 Адаптация атак, связанных с ИИ</span><p>Используя эти переведенные пары предложений, исследователи обучили модель, реплицирующую поведение целевой модели.</p></a>
<a class="relation-item" href="/studies/AML.CS0008/"><span class="relation-id">AML.CS0008</span><strong>Обход ProofPoint</strong><span class="relation-meta">Актор: Researchers at Silent Break Security / Тактика: AML.TA0001 Адаптация атак, связанных с ИИ</span><p>Исследователи использовали письма и собранные оценки как набор данных, на котором обучили рабочую копию модели ProofPoint. С помощью простой корреляции они определили, какая переменная оценки в целом отражает безопасность письма. В этом случае была выбрана переменная &#34;mlxlogscore&#34;, поскольку она была связана со spam, phish и core mlx, и ее использовали как метку. Каждое значение &#34;mlxlogscore&#34; обычно находилось в диапазоне от 1 до 999: чем выше оценка, тем безопаснее образец. Обучение выполнялось с использованием искусственной нейронной сети (ANN) и токенизации Bag of Words.</p></a>
</div>

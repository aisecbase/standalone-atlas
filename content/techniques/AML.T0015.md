---
atlas_id: AML.T0015
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут создавать состязательные данные, которые мешают ИИ-модели корректно определить содержимое данных, или создавать дипфейки, обманывающие ИИ-модель, рассчитанную на подлинные данные. Эта техника...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 7
modified_date: "2026-05-27"
platforms:
    - Predictive AI
procedure_count: 18
source_name: Evade AI Model
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0004
    - AML.TA0007
    - AML.TA0011
title: Обход ИИ-модели
url: /techniques/AML.T0015/
---

Злоумышленники могут [создавать состязательные данные](/techniques/AML.T0043), которые мешают ИИ-модели корректно определить содержимое данных, или [создавать дипфейки](/techniques/AML.T0088), обманывающие ИИ-модель, рассчитанную на подлинные данные.

Эта техника может использоваться для обхода последующей задачи, в которой применяется ИИ. Злоумышленник может обходить обнаружение вирусов или вредоносного ПО на основе ИИ либо сетевое сканирование для достижения целей традиционной кибератаки. Обход ИИ-модели с помощью дипфейков также может дать первичный доступ к системам, использующим биометрическую аутентификацию на основе ИИ.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0004/"><span class="relation-id">AML.TA0004</span><strong>Первичный доступ</strong></a>
<a class="relation-item" href="/tactics/AML.TA0007/"><span class="relation-id">AML.TA0007</span><strong>Уклонение от защиты</strong></a>
<a class="relation-item" href="/tactics/AML.TA0011/"><span class="relation-id">AML.TA0011</span><strong>Воздействие</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0003/"><span class="relation-id">AML.M0003</span><strong>Усиление устойчивости модели</strong><p>Усиленные модели труднее обходить.</p></a>
<a class="relation-item" href="/mitigations/AML.M0006/"><span class="relation-id">AML.M0006</span><strong>Использование ансамблевых методов</strong><p>Использование нескольких разных моделей повышает устойчивость к атакам.</p></a>
<a class="relation-item" href="/mitigations/AML.M0009/"><span class="relation-id">AML.M0009</span><strong>Использование мультимодальных сенсоров</strong><p>Использование разных сенсоров может затруднить злоумышленнику компрометацию системы и получение вредоносных результатов.</p></a>
<a class="relation-item" href="/mitigations/AML.M0010/"><span class="relation-id">AML.M0010</span><strong>Восстановление входных данных</strong><p>Предобработка входных данных модели может предотвратить прохождение вредоносных данных через пайплайн машинного обучения.</p></a>
<a class="relation-item" href="/mitigations/AML.M0015/"><span class="relation-id">AML.M0015</span><strong>Обнаружение состязательных входных данных</strong><p>Предотвращает внесение злоумышленником состязательных данных в систему.</p></a>
<a class="relation-item" href="/mitigations/AML.M0034/"><span class="relation-id">AML.M0034</span><strong>Обнаружение дипфейков</strong><p>Обнаружение дипфейков можно использовать для выявления и блокировки сгенерированного контента.</p></a>
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>Красная команда по ИИ</strong><p>Проведите репрезентативные цифровые и мультимодальные атаки обхода, а также атаки обхода в физическом домене. Используйте успешные атаки, чтобы повысить устойчивость модели и улучшить предобработку, обнаружение состязательных входных данных, контроль со стороны человека и мониторинг.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0000/"><span class="relation-id">AML.CS0000</span><strong>Обход детектора C&amp;C-трафика вредоносного ПО на основе глубокого обучения</strong><span class="relation-meta">Актор: Palo Alto Networks AI Research Team / Тактика: AML.TA0007 Уклонение от защиты</span><p>Используя созданные образцы, мы выполнили онлайн-обход модели обнаружения шпионского ПО на основе машинного обучения. Созданные пакеты были классифицированы как безвредные с уверенностью &gt; 80%. Эта оценка показывает, что злоумышленники способны обходить продвинутые техники обнаружения на основе машинного обучения, создавая образцы, которые ML-модель классифицирует неверно.</p></a>
<a class="relation-item" href="/studies/AML.CS0001/"><span class="relation-id">AML.CS0001</span><strong>Обход обнаружения DGA-доменов ботнетов</strong><span class="relation-meta">Актор: Palo Alto Networks AI Research Team / Тактика: AML.TA0007 Уклонение от защиты</span><p>Доменные имена, сгенерированные DGA и измененные с помощью этой техники, успешно обходят целевую модель обнаружения DGA, позволяя злоумышленнику продолжать связь со своими серверами [командования и управления](https://attack.mitre.org/tactics/TA0011/).</p></a>
<a class="relation-item" href="/studies/AML.CS0003/"><span class="relation-id">AML.CS0003</span><strong>Обход ИИ-детектора вредоносного ПО Cylance</strong><span class="relation-meta">Актор: Skylight Cyber / Тактика: AML.TA0007 Уклонение от защиты</span><p>Поскольку вторичная модель переопределяла основную, исследователи фактически смогли обойти ML-модель.</p></a>
<a class="relation-item" href="/studies/AML.CS0004/"><span class="relation-id">AML.CS0004</span><strong>Атака на систему распознавания лиц через подмену видеопотока камеры</strong><span class="relation-meta">Актор: Two individuals / Тактика: AML.TA0004 Первичный доступ</span><p>Злоумышленники успешно обошли систему распознавания лиц. Это позволило им выдать себя за жертву и подтвердить ее личность в налоговой системе.</p></a>
<a class="relation-item" href="/studies/AML.CS0005/"><span class="relation-id">AML.CS0005</span><strong>Атака на сервисы машинного перевода</strong><span class="relation-meta">Актор: Berkeley Artificial Intelligence Research / Тактика: AML.TA0011 Воздействие</span><p>Состязательные примеры использовались для обхода сервисов машинного перевода разными способами. Среди них были целевые подмены слов, нецензурный вывод и пропуски предложений.</p></a>
<a class="relation-item" href="/studies/AML.CS0008/"><span class="relation-id">AML.CS0008</span><strong>Обход ProofPoint</strong><span class="relation-meta">Актор: Researchers at Silent Break Security / Тактика: AML.TA0011 Воздействие</span><p>В итоге сведения, полученные с помощью &#34;офлайн&#34; прокси-модели, позволили исследователям создавать вредоносные письма, которые получали благоприятные оценки от реальной системы защиты электронной почты ProofPoint и тем самым обходили ее.</p></a>
<a class="relation-item" href="/studies/AML.CS0010/"><span class="relation-id">AML.CS0010</span><strong>Нарушение работы сервиса Microsoft Azure</strong><span class="relation-meta">Актор: Microsoft AI Red Team / Тактика: AML.TA0011 Воздействие</span><p>Команда провела онлайн-атаку обхода, повторно отправив состязательные примеры, и достигла своих целей.</p></a>
<a class="relation-item" href="/studies/AML.CS0011/"><span class="relation-id">AML.CS0011</span><strong>Обход ИИ на периферии Microsoft</strong><span class="relation-meta">Актор: Azure Red Team / Тактика: AML.TA0011 Воздействие</span><p>Красная команда подала на вход модели это незаметно измененное изображение и смогла обойти ML-модель, добившись ошибочной классификации.</p></a>
<a class="relation-item" href="/studies/AML.CS0012/"><span class="relation-id">AML.CS0012</span><strong>Обход системы идентификации лиц с помощью физических контрмер</strong><span class="relation-meta">Актор: MITRE AI Red Team / Тактика: AML.TA0011 Воздействие</span><p>Команда успешно обошла модель с помощью физической наклейки, вызвав целевые ошибочные классификации.</p></a>
<a class="relation-item" href="/studies/AML.CS0013/"><span class="relation-id">AML.CS0013</span><strong>Бэкдор-атака на модели глубокого обучения в мобильных приложениях</strong><span class="relation-meta">Актор: Yuanchun Li, Jiayi Hua, Haoyu Wang, Chunyang Chen, Yunxin Liu / Тактика: AML.TA0011 Воздействие</span><p>Предъявление визуального триггера приводит к обходу модели жертвы. Исследователи показали, что это можно использовать для обхода ML-моделей в нескольких критически важных для безопасности приложениях из Google Play.</p></a>
<a class="relation-item" href="/studies/AML.CS0014/"><span class="relation-id">AML.CS0014</span><strong>Сбивание с толку антивирусных нейронных сетей</strong><span class="relation-meta">Актор: Kaspersky ML Research Team / Тактика: AML.TA0007 Уклонение от защиты</span><p>Исследователи показали, что для большинства состязательных файлов антивирусная модель была успешно обойдена. На практике злоумышленник мог бы развернуть специально подготовленное вредоносное ПО и заражать системы, избегая обнаружения.</p></a>
<a class="relation-item" href="/studies/AML.CS0017/"><span class="relation-id">AML.CS0017</span><strong>Обход проверки личности ID.me</strong><span class="relation-meta">Актор: One individual / Тактика: AML.TA0004 Первичный доступ</span><p>Мужчина собрал украденные персональные данные, включая имена, даты рождения и номера социального страхования, и использовал их вместе со своей фотографией, на которой он был в парике, для получения поддельных водительских удостоверений. Он загрузил поддельные удостоверения вместе с селфи. Система проверки документов ID.me сопоставила селфи с фотографией в удостоверении, что позволило некоторым мошенническим заявкам пройти следующие этапы процесса обработки.</p></a>
<a class="relation-item" href="/studies/AML.CS0028/"><span class="relation-id">AML.CS0028</span><strong>Подмена ИИ-модели через атаку на цепочку поставок</strong><span class="relation-meta">Актор: Trend Micro Nebula Cloud Research Team / Тактика: AML.TA0011 Воздействие</span><p>После развертывания контейнерного образа злоумышленника модель может ошибочно классифицировать входные данные из-за внесенных им изменений.</p></a>
<a class="relation-item" href="/studies/AML.CS0032/"><span class="relation-id">AML.CS0032</span><strong>Попытка обхода ML-системы обнаружения фишинговых веб-страниц</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0007 Уклонение от защиты</span><p>Злоумышленникам удалось обойти модель визуального сходства, использовавшуюся для обнаружения имитации бренда. Однако другие компоненты системы обнаружения фишинга успешно выявили эти фишинговые сайты.</p></a>
<a class="relation-item" href="/studies/AML.CS0033/"><span class="relation-id">AML.CS0033</span><strong>Обход мобильной KYC-верификации с помощью дипфейк-изображения в реальном времени</strong><span class="relation-meta">Актор: iProov Red Team / Тактика: AML.TA0004 Первичный доступ</span><p>Исследователи транслируют дипфейк-видеопоток через OBS и используют приложение Virtual Camera, чтобы заменить стандартную камеру этим потоком. Это позволяет успешно обойти систему распознавания лица и пройти аутентификацию под личностью жертвы.</p></a>
<a class="relation-item" href="/studies/AML.CS0034/"><span class="relation-id">AML.CS0034</span><strong>ProKYC: дипфейк-инструмент для атак с мошенническим созданием аккаунтов</strong><span class="relation-meta">Актор: ProKYC, cybercriminal group / Тактика: AML.TA0004 Первичный доступ</span><p>Злоумышленник использовал ProKYC, чтобы заменить видеопоток с камеры дипфейк-видео с селфи. Это позволило успешно обойти KYC-проверку и пройти аутентификацию под поддельной личностью.</p></a>
<a class="relation-item" href="/studies/AML.CS0043/"><span class="relation-id">AML.CS0043</span><strong>Прототип вредоносного ПО со встроенной промпт-инъекцией</strong><span class="relation-meta">Актор: Unknown Threat Actor / Тактика: AML.TA0007 Уклонение от защиты</span><p>LLM-инструмент обнаружения или анализа вредоносного ПО мог быть сманипулирован так, чтобы не классифицировать бинарный файл Skynet как вредоносный. Примечание: промпт-инъекция не сработала против LLM, которые тестировали Check Point Research.</p></a>
<a class="relation-item" href="/studies/AML.CS0058/"><span class="relation-id">AML.CS0058</span><strong>Извлечение ИИ-моделей из Google Photos</strong><span class="relation-meta">Актор: Skyld / Тактика: AML.TA0011 Воздействие</span><p>Состязательные данные могли использоваться для обхода или иного ухудшения работы моделей Google Photos, включая обнаружение лиц и объектов.</p></a>
</div>

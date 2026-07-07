---
atlas_id: AML.TA0000
atlas_type: tactic
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленник пытается получить некоторый уровень доступа к ИИ-модели. Доступ к ИИ-модели включает техники, использующие разные виды доступа к модели. Такой доступ может применяться злоумышленником для получения...
generated: true
generated_by: atlasgen
modified_date: "2025-10-13"
procedure_count: 30
source_name: AI Model Access
technique_count: 4
title: Доступ к ИИ-модели
url: /tactics/AML.TA0000/
---

Злоумышленник пытается получить некоторый уровень доступа к ИИ-модели.

Доступ к ИИ-модели включает техники, использующие разные виды доступа к модели. Такой доступ может применяться злоумышленником для получения информации, разработки атак, а также как способ подачи данных на вход модели. Уровень доступа может варьироваться от полного знания внутреннего устройства модели до доступа к физической среде, где собираются данные, используемые ИИ-моделью. На разных этапах атаки злоумышленник может использовать разные уровни доступа к модели — от подготовки атаки до воздействия на целевую систему.

Доступ к ИИ-модели может требовать доступа к системе, в которой размещена модель; модель может быть публично доступна через API; либо доступ может быть получен косвенно через взаимодействие с продуктом или сервисом, который использует ИИ в своих процессах.


## Техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0040/"><span class="relation-id">AML.T0040</span><strong>Доступ к API инференса ИИ-модели</strong></a>
<a class="relation-item" href="/techniques/AML.T0041/"><span class="relation-id">AML.T0041</span><strong>Доступ к физической среде</strong></a>
<a class="relation-item" href="/techniques/AML.T0044/"><span class="relation-id">AML.T0044</span><strong>Полный доступ к ИИ-модели</strong></a>
<a class="relation-item" href="/techniques/AML.T0047/"><span class="relation-id">AML.T0047</span><strong>Продукт или сервис с поддержкой ИИ</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0003/"><span class="relation-id">AML.CS0003</span><strong>Обход ИИ-детектора вредоносного ПО Cylance</strong><span class="relation-meta">Актор: Skylight Cyber / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Исследователи имели доступ к ПО Cylance для обнаружения вредоносного ПО с поддержкой ИИ.</p></a>
<a class="relation-item" href="/studies/AML.CS0004/"><span class="relation-id">AML.CS0004</span><strong>Атака на систему распознавания лиц через подмену видеопотока камеры</strong><span class="relation-meta">Актор: Two individuals / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Злоумышленники использовали приложение виртуальной камеры, чтобы передать сгенерированное видео в сервис распознавания лиц на основе машинного обучения, применявшийся для проверки пользователей.</p></a>
<a class="relation-item" href="/studies/AML.CS0005/"><span class="relation-id">AML.CS0005</span><strong>Атака на сервисы машинного перевода</strong><span class="relation-meta">Актор: Berkeley Artificial Intelligence Research / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Исследователи использовали публично доступное приложение не по назначению: отправляли запросы к модели и получали машинно переведенные пары предложений в качестве обучающих данных.</p></a>
<a class="relation-item" href="/studies/AML.CS0008/"><span class="relation-id">AML.CS0008</span><strong>Обход ProofPoint</strong><span class="relation-meta">Актор: Researchers at Silent Break Security / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Исследователи отправили через систему множество писем, чтобы собрать выходные данные модели из заголовков.</p></a>
<a class="relation-item" href="/studies/AML.CS0009/"><span class="relation-id">AML.CS0009</span><strong>Отравление Tay</strong><span class="relation-meta">Актор: 4chan Users / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Злоумышленники могли взаимодействовать с Tay через сообщения в Twitter.</p></a>
<a class="relation-item" href="/studies/AML.CS0010/"><span class="relation-id">AML.CS0010</span><strong>Нарушение работы сервиса Microsoft Azure</strong><span class="relation-meta">Актор: Microsoft AI Red Team / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Команда использовала открытый API для доступа к целевой модели.</p></a>
<a class="relation-item" href="/studies/AML.CS0011/"><span class="relation-id">AML.CS0011</span><strong>Обход ИИ на периферии Microsoft</strong><span class="relation-meta">Актор: Azure Red Team / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Используя общедоступную версию ML-модели, команда начала отправлять запросы и анализировать ответы, то есть результаты инференса модели.</p></a>
<a class="relation-item" href="/studies/AML.CS0012/"><span class="relation-id">AML.CS0012</span><strong>Обход системы идентификации лиц с помощью физических контрмер</strong><span class="relation-meta">Актор: MITRE AI Red Team / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Команда получила доступ к API инференса целевой модели.</p></a>
<a class="relation-item" href="/studies/AML.CS0012/"><span class="relation-id">AML.CS0012</span><strong>Обход системы идентификации лиц с помощью физических контрмер</strong><span class="relation-meta">Актор: MITRE AI Red Team / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Команда разместила подготовленную наклейку в физической среде, чтобы вызвать сбои в системе идентификации лиц.</p></a>
<a class="relation-item" href="/studies/AML.CS0013/"><span class="relation-id">AML.CS0013</span><strong>Бэкдор-атака на модели глубокого обучения в мобильных приложениях</strong><span class="relation-meta">Актор: Yuanchun Li, Jiayi Hua, Haoyu Wang, Chunyang Chen, Yunxin Liu / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>На этапе инференса для запуска атаки требуется только доступ к физической среде.</p></a>
<a class="relation-item" href="/studies/AML.CS0013/"><span class="relation-id">AML.CS0013</span><strong>Бэкдор-атака на модели глубокого обучения в мобильных приложениях</strong><span class="relation-meta">Актор: Yuanchun Li, Jiayi Hua, Haoyu Wang, Chunyang Chen, Yunxin Liu / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Это дало исследователям полный доступ к ML-модели, хотя и в скомпилированном бинарном виде.</p></a>
<a class="relation-item" href="/studies/AML.CS0014/"><span class="relation-id">AML.CS0014</span><strong>Сбивание с толку антивирусных нейронных сетей</strong><span class="relation-meta">Актор: Kaspersky ML Research Team / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>На протяжении кейса исследователи использовали доступ к целевому антивирусному продукту на основе ML. Продукт сканирует файлы на системе пользователя, локально извлекает признаки, а затем отправляет их в облачный ML-детектор вредоносного ПО для классификации. Поэтому у исследователей был только доступ к самому детектору в режиме чёрного ящика, но из экстрактора признаков они могли получить ценную информацию для построения атаки.</p></a>
</div>


Показано 12 из 30 примеров.

---
atlas_id: AML.T0043
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Состязательные данные — это входные данные для ИИ-модели, изменённые таким образом, чтобы вызвать в целевой модели желаемый злоумышленником эффект. Возможные эффекты варьируются от ошибочной классификации и пропусков...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 8
modified_date: "2026-05-27"
platforms:
    - Predictive AI
procedure_count: 1
source_name: Craft Adversarial Data
subtechnique_count: 5
subtechnique_of: ""
tactics:
    - AML.TA0001
title: Создание состязательных данных
url: /techniques/AML.T0043/
---

Состязательные данные — это входные данные для ИИ-модели, изменённые таким образом, чтобы вызвать в целевой модели желаемый злоумышленником эффект.

Возможные эффекты варьируются от ошибочной классификации и пропусков обнаружения до максимизации энергопотребления.

Обычно величину или область модификации ограничивают так, чтобы человек продолжал воспринимать данные как неизменённые. Однако в зависимости от желаемого злоумышленником эффекта незаметность изменений для человека не всегда имеет значение.

Например, состязательным входом для задачи классификации изображений служит изображение, которое ИИ-модель классифицирует неверно, хотя человек по-прежнему распознаёт в нём объект правильного класса.

В зависимости от знаний злоумышленника о целевой модели и уровня доступа к ней он может использовать для разработки состязательного примера различные классы алгоритмов, например [оптимизацию в режиме белого ящика](/techniques/AML.T0043.000), [оптимизацию в режиме чёрного ящика](/techniques/AML.T0043.001), [перенос на модель чёрного ящика](/techniques/AML.T0043.002) или [ручную модификацию](/techniques/AML.T0043.003).

Если у злоумышленника есть доступ к модели в режиме белого ящика или через API инференса, он может выполнить [проверку атаки](/techniques/AML.T0042), чтобы подтвердить работоспособность своего подхода.

Это позволяет злоумышленнику убедиться, что его атака эффективна в реальной среде, где её могут заметить.

Затем он может применить атаку позднее для достижения своих целей.

Злоумышленник может оптимизировать состязательные примеры для [обхода ИИ-модели](/techniques/AML.T0015) или [нарушения целостности ИИ-модели](/techniques/AML.T0031).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0001/"><span class="relation-id">AML.TA0001</span><strong>Подготовка атаки на ИИ</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0043.000/"><span class="relation-id">AML.T0043.000</span><strong>Оптимизация в режиме белого ящика</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.001/"><span class="relation-id">AML.T0043.001</span><strong>Оптимизация в режиме чёрного ящика</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.002/"><span class="relation-id">AML.T0043.002</span><strong>Перенос на модель чёрного ящика</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.003/"><span class="relation-id">AML.T0043.003</span><strong>Ручная модификация</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.004/"><span class="relation-id">AML.T0043.004</span><strong>Добавление бэкдор-триггера</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0002/"><span class="relation-id">AML.M0002</span><strong>Обфускация выходных данных предиктивного ИИ</strong><p>Обфускация выходных данных модели снижает способность злоумышленника генерировать эффективные состязательные данные.</p></a>
<a class="relation-item" href="/mitigations/AML.M0003/"><span class="relation-id">AML.M0003</span><strong>Повышение устойчивости моделей предиктивного ИИ</strong><p>Модели с повышенной устойчивостью лучше противостоят состязательным входным данным.</p></a>
<a class="relation-item" href="/mitigations/AML.M0004/"><span class="relation-id">AML.M0004</span><strong>Ограничение объёма и частоты запросов к ИИ-сервису</strong><p>Ограничивайте объём запросов к модели, чтобы помешать злоумышленнику создавать состязательные входные данные или замедлить их создание.</p></a>
<a class="relation-item" href="/mitigations/AML.M0006/"><span class="relation-id">AML.M0006</span><strong>Ансамбли моделей предиктивного ИИ</strong><p>Использование ансамбля моделей усложняет создание эффективных состязательных данных и повышает общую устойчивость.</p></a>
<a class="relation-item" href="/mitigations/AML.M0008/"><span class="relation-id">AML.M0008</span><strong>Валидация ИИ-модели</strong><p>Проверка ИИ-модели на состязательных данных помогает убедиться, что модель работает как задумано и устойчива к состязательным входным данным.</p></a>
<a class="relation-item" href="/mitigations/AML.M0010/"><span class="relation-id">AML.M0010</span><strong>Восстановление входных данных предиктивного ИИ</strong><p>Восстановление входных данных может помогать исправлять состязательные входные данные.</p></a>
<a class="relation-item" href="/mitigations/AML.M0015/"><span class="relation-id">AML.M0015</span><strong>Обнаружение состязательных входных данных для предиктивного ИИ</strong><p>Встраивайте обнаружение состязательных входных данных, чтобы блокировать вредоносные входные данные во время инференса.</p></a>
<a class="relation-item" href="/mitigations/AML.M0019/"><span class="relation-id">AML.M0019</span><strong>Контроль доступа к ИИ-моделям и данным в продакшене</strong><p>Средства контроля доступа к API модели могут ограничить необходимый злоумышленнику доступ для создания состязательных данных.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0002/"><span class="relation-id">AML.CS0002</span><strong>Отравление VirusTotal</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Злоумышленник использовал образец вредоносного ПО из распространенного семейства программ-вымогателей как исходную точку для создания мутированных вариантов.</p></a>
</div>

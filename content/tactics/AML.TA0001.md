---
atlas_id: AML.TA0001
atlas_type: tactic
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленник использует свои знания о целевой системе и доступ к ней, чтобы адаптировать атаку. Подготовка атаки на ИИ включает техники, которые злоумышленники используют для подготовки атаки на целевую ИИ-модель....
generated: true
generated_by: atlasgen
modified_date: "2026-08-31"
procedure_count: 76
source_name: AI Attack Adaptation
technique_count: 39
title: Подготовка атаки на ИИ
url: /tactics/AML.TA0001/
---

Злоумышленник использует свои знания о целевой системе и доступ к ней, чтобы адаптировать атаку.

Подготовка атаки на ИИ включает техники, которые злоумышленники используют для подготовки атаки на целевую ИИ-модель.
Такие техники могут включать обучение прокси-моделей, отравление целевой модели и создание состязательных данных для подачи в целевую модель.
Некоторые из этих техник могут выполняться в офлайн-режиме, поэтому им трудно противодействовать.
Эти техники часто используются для достижения конечной цели злоумышленника.


## Техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0005/"><span class="relation-id">AML.T0005</span><strong>Создание прокси-модели ИИ</strong></a>
<a class="relation-item" href="/techniques/AML.T0005.000/"><span class="relation-id">AML.T0005.000</span><strong>Обучение прокси-модели на собранных ИИ-артефактах</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0005.000/"><span class="relation-id">AML.T0005.000</span><strong>Обучение прокси-модели на собранных ИИ-артефактах</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0005.001/"><span class="relation-id">AML.T0005.001</span><strong>Обучение прокси-модели через репликацию</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0005.001/"><span class="relation-id">AML.T0005.001</span><strong>Обучение прокси-модели через репликацию</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0005.002/"><span class="relation-id">AML.T0005.002</span><strong>Использование предварительно обученной модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0005.002/"><span class="relation-id">AML.T0005.002</span><strong>Использование предварительно обученной модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018/"><span class="relation-id">AML.T0018</span><strong>Манипуляция ИИ-моделью</strong></a>
<a class="relation-item" href="/techniques/AML.T0018.000/"><span class="relation-id">AML.T0018.000</span><strong>Отравление ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.000/"><span class="relation-id">AML.T0018.000</span><strong>Отравление ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.001/"><span class="relation-id">AML.T0018.001</span><strong>Изменение архитектуры ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.001/"><span class="relation-id">AML.T0018.001</span><strong>Изменение архитектуры ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.002/"><span class="relation-id">AML.T0018.002</span><strong>Встраивание вредоносного ПО</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.002/"><span class="relation-id">AML.T0018.002</span><strong>Встраивание вредоносного ПО</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.003/"><span class="relation-id">AML.T0018.003</span><strong>Изменение логики формирования промпта</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.003/"><span class="relation-id">AML.T0018.003</span><strong>Изменение логики формирования промпта</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0042/"><span class="relation-id">AML.T0042</span><strong>Проверка атаки</strong></a>
<a class="relation-item" href="/techniques/AML.T0043/"><span class="relation-id">AML.T0043</span><strong>Создание состязательных данных</strong></a>
<a class="relation-item" href="/techniques/AML.T0043.000/"><span class="relation-id">AML.T0043.000</span><strong>Оптимизация в режиме белого ящика</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.000/"><span class="relation-id">AML.T0043.000</span><strong>Оптимизация в режиме белого ящика</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.001/"><span class="relation-id">AML.T0043.001</span><strong>Оптимизация в режиме чёрного ящика</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.001/"><span class="relation-id">AML.T0043.001</span><strong>Оптимизация в режиме чёрного ящика</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.002/"><span class="relation-id">AML.T0043.002</span><strong>Перенос на модель чёрного ящика</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.002/"><span class="relation-id">AML.T0043.002</span><strong>Перенос на модель чёрного ящика</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.003/"><span class="relation-id">AML.T0043.003</span><strong>Ручная модификация</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.003/"><span class="relation-id">AML.T0043.003</span><strong>Ручная модификация</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.004/"><span class="relation-id">AML.T0043.004</span><strong>Добавление бэкдор-триггера</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0043.004/"><span class="relation-id">AML.T0043.004</span><strong>Добавление бэкдор-триггера</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0065/"><span class="relation-id">AML.T0065</span><strong>Создание промптов для LLM</strong></a>
<a class="relation-item" href="/techniques/AML.T0066/"><span class="relation-id">AML.T0066</span><strong>Подготовка содержимого для извлечения</strong></a>
<a class="relation-item" href="/techniques/AML.T0088/"><span class="relation-id">AML.T0088</span><strong>Создание дипфейков</strong></a>
<a class="relation-item" href="/techniques/AML.T0102/"><span class="relation-id">AML.T0102</span><strong>Генерация вредоносных команд</strong></a>
<a class="relation-item" href="/techniques/AML.T0117/"><span class="relation-id">AML.T0117</span><strong>Автономная адаптация пути атаки</strong></a>
<a class="relation-item" href="/techniques/AML.T0118/"><span class="relation-id">AML.T0118</span><strong>Коммуникация автономных ИИ-агентов</strong></a>
<a class="relation-item" href="/techniques/AML.T0118.000/"><span class="relation-id">AML.T0118.000</span><strong>Communication via Shared Artifacts</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0118.000/"><span class="relation-id">AML.T0118.000</span><strong>Communication via Shared Artifacts</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0118.001/"><span class="relation-id">AML.T0118.001</span><strong>Direct Agent Communication</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0118.001/"><span class="relation-id">AML.T0118.001</span><strong>Direct Agent Communication</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0124/"><span class="relation-id">AML.T0124</span><strong>Autonomous Attack Orchestration</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0000/"><span class="relation-id">AML.CS0000</span><strong>Обход детектора C&amp;C-трафика вредоносного ПО на основе глубокого обучения</strong><span class="relation-meta">Актор: Palo Alto Networks AI Research Team / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Мы обучили модель на наборе данных HTTP-трафика, чтобы использовать ее как прокси для целевой модели. Оценка показала среднюю долю истинно положительных срабатываний около 99% и среднюю долю ложноположительных срабатываний около 0,01%. При проверке модели заголовок HTTP-пакета из известных образцов C&amp;C-трафика вредоносного ПО был классифицирован как вредоносный с высокой уверенностью (&gt; 99%).</p></a>
<a class="relation-item" href="/studies/AML.CS0000/"><span class="relation-id">AML.CS0000</span><strong>Обход детектора C&amp;C-трафика вредоносного ПО на основе глубокого обучения</strong><span class="relation-meta">Актор: Palo Alto Networks AI Research Team / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Мы отправляли в модель наши состязательные примеры и корректировали их, пока модель не удалось обойти.</p></a>
<a class="relation-item" href="/studies/AML.CS0000/"><span class="relation-id">AML.CS0000</span><strong>Обход детектора C&amp;C-трафика вредоносного ПО на основе глубокого обучения</strong><span class="relation-meta">Актор: Palo Alto Networks AI Research Team / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Мы создали образцы для обхода, удалив из заголовка пакета поля, которые обычно не используются для C&amp;C-коммуникации, например `cache-control`, `connection` и т. д.</p></a>
<a class="relation-item" href="/studies/AML.CS0001/"><span class="relation-id">AML.CS0001</span><strong>Обход обнаружения DGA-доменов ботнетов</strong><span class="relation-meta">Актор: Palo Alto Networks AI Research Team / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Результаты эксперимента показали, что доля обнаружения всех 16 семейств ботнетных DGA падает ниже 25% после однократной вставки всего одной строки в доменные имена, сгенерированные DGA.</p></a>
<a class="relation-item" href="/studies/AML.CS0001/"><span class="relation-id">AML.CS0001</span><strong>Обход обнаружения DGA-доменов ботнетов</strong><span class="relation-meta">Актор: Palo Alto Networks AI Research Team / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Исследователи использовали технику мутации для генерации доменных имен, обходящих обнаружение.</p></a>
<a class="relation-item" href="/studies/AML.CS0002/"><span class="relation-id">AML.CS0002</span><strong>Отравление VirusTotal</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Злоумышленник использовал образец вредоносного ПО из распространенного семейства программ-вымогателей как исходную точку для создания мутированных вариантов.</p></a>
<a class="relation-item" href="/studies/AML.CS0003/"><span class="relation-id">AML.CS0003</span><strong>Обход ИИ-детектора вредоносного ПО Cylance</strong><span class="relation-meta">Актор: Skylight Cyber / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Используя эти знания, исследователи объединили атрибуты заведомо легитимных файлов с вредоносным ПО, чтобы вручную создать состязательные образцы вредоносного ПО.</p></a>
<a class="relation-item" href="/studies/AML.CS0005/"><span class="relation-id">AML.CS0005</span><strong>Атака на сервисы машинного перевода</strong><span class="relation-meta">Актор: Berkeley Artificial Intelligence Research / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Используя эти переведенные пары предложений, исследователи обучили модель, реплицирующую поведение целевой модели.</p></a>
<a class="relation-item" href="/studies/AML.CS0005/"><span class="relation-id">AML.CS0005</span><strong>Атака на сервисы машинного перевода</strong><span class="relation-meta">Актор: Berkeley Artificial Intelligence Research / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Реплицированные модели использовались для генерации состязательных примеров, которые успешно срабатывали на сервисах машинного перевода с закрытой внутренней логикой.</p></a>
<a class="relation-item" href="/studies/AML.CS0007/"><span class="relation-id">AML.CS0007</span><strong>Репликация модели GPT-2</strong><span class="relation-meta">Актор: Researchers at Brown University / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Исследователи изменили функцию потерь Grover так, чтобы она соответствовала функции потерь GPT-2, а затем обучили модель на подготовленном ими наборе данных, используя исходные гиперпараметры Grover. Полученная модель воспроизводила поведение GPT-2 и показывала сопоставимое качество на большинстве наборов данных. Злоумышленник, повторивший действия исследователей, мог бы затем использовать такую копию GPT-2 во вредоносных целях.</p></a>
<a class="relation-item" href="/studies/AML.CS0008/"><span class="relation-id">AML.CS0008</span><strong>Обход ProofPoint</strong><span class="relation-meta">Актор: Researchers at Silent Break Security / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Исследователи использовали письма и собранные оценки как набор данных, на котором обучили рабочую копию модели ProofPoint. С помощью простой корреляции они определили, какая переменная оценки в целом отражает безопасность письма. В этом случае была выбрана переменная &#34;mlxlogscore&#34;, поскольку она была связана со spam, phish и core mlx, и ее использовали как метку. Каждое значение &#34;mlxlogscore&#34; обычно находилось в диапазоне от 1 до 999: чем выше оценка, тем безопаснее образец. Обучение выполнялось с использованием искусственной нейронной сети (ANN) и токенизации Bag of Words.</p></a>
<a class="relation-item" href="/studies/AML.CS0008/"><span class="relation-id">AML.CS0008</span><strong>Обход ProofPoint</strong><span class="relation-meta">Актор: Researchers at Silent Break Security / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Затем исследователи ML алгоритмически нашли в этой &#34;офлайн&#34; прокси-модели образцы, которые помогли получить нужное представление о ее поведении и влиятельных переменных. Примеры образцов с хорошими оценками: &#34;calculation&#34;, &#34;asset&#34; и &#34;tyson&#34;. Примеры образцов с плохими оценками: &#34;software&#34;, &#34;99&#34; и &#34;unsub&#34;.</p></a>
</div>


Показано 12 из 76 примеров.

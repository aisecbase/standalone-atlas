---
atlas_id: AML.TA0006
atlas_type: tactic
attack_ref_id: TA0003
attack_ref_url: https://attack.mitre.org/tactics/TA0003/
created_date: "2022-01-24"
description: Злоумышленник пытается сохранить закрепление с помощью ИИ-артефактов или программного обеспечения. Закрепление включает техники, которые злоумышленники используют, чтобы сохранять доступ к системам после перезагрузок,...
generated: true
generated_by: atlasgen
modified_date: "2025-04-09"
procedure_count: 24
source_name: Persistence
technique_count: 29
title: Закрепление
url: /tactics/AML.TA0006/
---

Злоумышленник пытается сохранить закрепление с помощью ИИ-артефактов или программного обеспечения.

Закрепление включает техники, которые злоумышленники используют, чтобы сохранять доступ к системам после перезагрузок, изменения учетных данных и других прерываний, способных лишить их доступа. Техники закрепления часто предполагают оставление в системе модифицированных ML-артефактов, таких как отравленные обучающие данные или измененные ИИ-модели.


## Техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0018/"><span class="relation-id">AML.T0018</span><strong>Манипуляция ИИ-моделью</strong></a>
<a class="relation-item" href="/techniques/AML.T0018.000/"><span class="relation-id">AML.T0018.000</span><strong>Отравление ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.000/"><span class="relation-id">AML.T0018.000</span><strong>Отравление ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.001/"><span class="relation-id">AML.T0018.001</span><strong>Изменение архитектуры ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.001/"><span class="relation-id">AML.T0018.001</span><strong>Изменение архитектуры ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.002/"><span class="relation-id">AML.T0018.002</span><strong>Встраивание вредоносного ПО</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.002/"><span class="relation-id">AML.T0018.002</span><strong>Встраивание вредоносного ПО</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.003/"><span class="relation-id">AML.T0018.003</span><strong>Изменение логики формирования промпта</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.003/"><span class="relation-id">AML.T0018.003</span><strong>Изменение логики формирования промпта</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0020/"><span class="relation-id">AML.T0020</span><strong>Отравление обучающих данных</strong></a>
<a class="relation-item" href="/techniques/AML.T0061/"><span class="relation-id">AML.T0061</span><strong>Саморепликация промпта LLM</strong></a>
<a class="relation-item" href="/techniques/AML.T0070/"><span class="relation-id">AML.T0070</span><strong>Отравление RAG</strong></a>
<a class="relation-item" href="/techniques/AML.T0080/"><span class="relation-id">AML.T0080</span><strong>Отравление контекста ИИ-агента</strong></a>
<a class="relation-item" href="/techniques/AML.T0080.000/"><span class="relation-id">AML.T0080.000</span><strong>Память</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0080.000/"><span class="relation-id">AML.T0080.000</span><strong>Память</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0080.001/"><span class="relation-id">AML.T0080.001</span><strong>Цепочка сообщений</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0080.001/"><span class="relation-id">AML.T0080.001</span><strong>Цепочка сообщений</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0081/"><span class="relation-id">AML.T0081</span><strong>Изменение конфигурации ИИ-агента</strong></a>
<a class="relation-item" href="/techniques/AML.T0093/"><span class="relation-id">AML.T0093</span><strong>Внедрение промпта через публичное приложение</strong></a>
<a class="relation-item" href="/techniques/AML.T0099/"><span class="relation-id">AML.T0099</span><strong>Отравление данных инструмента ИИ-агента</strong></a>
<a class="relation-item" href="/techniques/AML.T0110/"><span class="relation-id">AML.T0110</span><strong>Отравление инструмента ИИ-агента</strong></a>
<a class="relation-item" href="/techniques/AML.T0110.000/"><span class="relation-id">AML.T0110.000</span><strong>Определение и инструкции</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0110.000/"><span class="relation-id">AML.T0110.000</span><strong>Определение и инструкции</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0110.001/"><span class="relation-id">AML.T0110.001</span><strong>Реализация</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0110.001/"><span class="relation-id">AML.T0110.001</span><strong>Реализация</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0110.002/"><span class="relation-id">AML.T0110.002</span><strong>Ответ во время выполнения</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0110.002/"><span class="relation-id">AML.T0110.002</span><strong>Ответ во время выполнения</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0121/"><span class="relation-id">AML.T0121</span><strong>AI Agent Environment Reconstruction</strong></a>
<a class="relation-item" href="/techniques/AML.T0125/"><span class="relation-id">AML.T0125</span><strong>Create Account</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0002/"><span class="relation-id">AML.CS0002</span><strong>Отравление VirusTotal</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0006 Закрепление</span><p>Несколько вендоров начали классифицировать файлы как относящиеся к этому семейству программ-вымогателей, хотя большинство из них не запускались. Мутированные образцы отравили набор данных, который модели машинного обучения используют для выявления и классификации этого семейства программ-вымогателей.</p></a>
<a class="relation-item" href="/studies/AML.CS0009/"><span class="relation-id">AML.CS0009</span><strong>Отравление Tay</strong><span class="relation-meta">Актор: 4chan Users / Тактика: AML.TA0006 Закрепление</span><p>Многократно взаимодействуя с Tay с использованием расистской и оскорбительной лексики, злоумышленники смогли сместить набор данных Tay в сторону такой же лексики. Для этого они использовали функцию &#34;repeat after me&#34; — команду, которая заставляла Tay повторять все, что ей говорили.</p></a>
<a class="relation-item" href="/studies/AML.CS0013/"><span class="relation-id">AML.CS0013</span><strong>Бэкдор-атака на модели глубокого обучения в мобильных приложениях</strong><span class="relation-meta">Актор: Yuanchun Li, Jiayi Hua, Haoyu Wang, Chunyang Chen, Yunxin Liu / Тактика: AML.TA0006 Закрепление</span><p>Исследователи отравили модель жертвы, внедрив нейронную полезную нагрузку в скомпилированные модели через прямое изменение вычислительного графа. Затем исследователи снова упаковали отравленную модель в APK-файл.</p></a>
<a class="relation-item" href="/studies/AML.CS0024/"><span class="relation-id">AML.CS0024</span><strong>Червь Morris II: атака на основе RAG</strong><span class="relation-meta">Актор: Stav Cohen, Ron Bitton, Ben Nassi / Тактика: AML.TA0006 Закрепление</span><p>Самореплицирующаяся часть промпта заставляет сгенерированный ответ содержать вредоносный промпт, благодаря чему червь может распространяться дальше.</p></a>
<a class="relation-item" href="/studies/AML.CS0025/"><span class="relation-id">AML.CS0025</span><strong>Отравление крупномасштабных веб-датасетов: атака split-view</strong><span class="relation-meta">Актор: Researchers from Google Deepmind, ETH Zurich, NVIDIA, Robust Intelligence, and Google / Тактика: AML.TA0006 Закрепление</span><p>Злоумышленник может подготовить отравленные обучающие данные, чтобы заменить фрагменты датасета, ставшие недоступными из-за истечения регистрации соответствующих доменов.</p></a>
<a class="relation-item" href="/studies/AML.CS0026/"><span class="relation-id">AML.CS0026</span><strong>Перехват финансовой транзакции с использованием M365 Copilot в роли инсайдера</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0006 Закрепление</span><p>Исследователи добились закрепления в системе жертвы, поскольку вредоносный промпт выполнялся каждый раз, когда извлекалась отравленная RAG-запись с поддельными банковскими реквизитами.</p></a>
<a class="relation-item" href="/studies/AML.CS0028/"><span class="relation-id">AML.CS0028</span><strong>Подмена ИИ-модели через атаку на цепочку поставок</strong><span class="relation-meta">Актор: Trend Micro Nebula Cloud Research Team / Тактика: AML.TA0006 Закрепление</span><p>Имея полный доступ к весам модели, злоумышленник мог изменить их, чтобы вызывать ошибочные классификации или иным образом ухудшать качество работы модели.</p></a>
<a class="relation-item" href="/studies/AML.CS0028/"><span class="relation-id">AML.CS0028</span><strong>Подмена ИИ-модели через атаку на цепочку поставок</strong><span class="relation-meta">Актор: Trend Micro Nebula Cloud Research Team / Тактика: AML.TA0006 Закрепление</span><p>Имея полный доступ к модели, злоумышленник мог изменить ее архитектуру, чтобы поменять поведение модели.</p></a>
<a class="relation-item" href="/studies/AML.CS0035/"><span class="relation-id">AML.CS0035</span><strong>Эксфильтрация данных из Slack AI через косвенную промпт-инъекцию</strong><span class="relation-meta">Актор: PromptArmor / Тактика: AML.TA0006 Закрепление</span><p>Исследователь создает публичный канал Slack и отправляет в него вредоносное содержимое: текст для извлечения и промпт. Поскольку Slack AI индексирует сообщения в публичных каналах, вредоносное сообщение добавляется в его RAG-базу данных.</p></a>
<a class="relation-item" href="/studies/AML.CS0036/"><span class="relation-id">AML.CS0036</span><strong>AIKatz: атака на десктопные LLM-приложения</strong><span class="relation-meta">Актор: Lumia Security / Тактика: AML.TA0006 Закрепление</span><p>Затем злоумышленник мог создавать вредоносные промпты, манипулирующие памятью LLM для достижения устойчивого эффекта. Любое изменение в памяти также распространялось бы на новые цепочки сообщений.</p></a>
<a class="relation-item" href="/studies/AML.CS0036/"><span class="relation-id">AML.CS0036</span><strong>AIKatz: атака на десктопные LLM-приложения</strong><span class="relation-meta">Актор: Lumia Security / Тактика: AML.TA0006 Закрепление</span><p>Злоумышленник мог создавать вредоносные промпты, манипулирующие контекстом цепочки сообщений; этот эффект сохранялся бы до конца такой цепочки.</p></a>
<a class="relation-item" href="/studies/AML.CS0040/"><span class="relation-id">AML.CS0040</span><strong>Взлом памяти ChatGPT с помощью промпт-инъекции</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0006 Закрепление</span><p>Промпт добавлял новые воспоминания и изменял поведение ChatGPT. Окно чата показывало, что память задана, хотя проверки или вмешательства со стороны человека не было. Все будущие сессии чата будут использовать отравленное хранилище памяти.</p></a>
</div>


Показано 12 из 24 примеров.

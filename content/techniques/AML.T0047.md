---
atlas_id: AML.T0047
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут использовать продукт или сервис, в котором искусственный интеллект применяется как часть внутренней реализации, чтобы получить доступ к лежащей в его основе ИИ-модели. Такой непрямой доступ к...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 16
source_name: AI-Enabled Product or Service
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0000
title: Продукт или сервис с поддержкой ИИ
url: /techniques/AML.T0047/
---

Злоумышленники могут использовать продукт или сервис, в котором искусственный интеллект применяется как часть внутренней реализации, чтобы получить доступ к лежащей в его основе ИИ-модели. Такой непрямой доступ к модели может раскрывать сведения об ИИ-модели или результатах ее инференса через журналы или метаданные.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0000/"><span class="relation-id">AML.TA0000</span><strong>Доступ к ИИ-модели</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0024/"><span class="relation-id">AML.M0024</span><strong>Логирование телеметрии ИИ</strong><p>Логирование телеметрии может помочь выявить отправку чувствительной информации о модели злоумышленнику.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0003/"><span class="relation-id">AML.CS0003</span><strong>Обход ИИ-детектора вредоносного ПО Cylance</strong><span class="relation-meta">Актор: Skylight Cyber / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Исследователи имели доступ к ПО Cylance для обнаружения вредоносного ПО с поддержкой ИИ.</p></a>
<a class="relation-item" href="/studies/AML.CS0004/"><span class="relation-id">AML.CS0004</span><strong>Атака на систему распознавания лиц через подмену видеопотока камеры</strong><span class="relation-meta">Актор: Two individuals / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Злоумышленники использовали приложение виртуальной камеры, чтобы передать сгенерированное видео в сервис распознавания лиц на основе машинного обучения, применявшийся для проверки пользователей.</p></a>
<a class="relation-item" href="/studies/AML.CS0008/"><span class="relation-id">AML.CS0008</span><strong>Обход ProofPoint</strong><span class="relation-meta">Актор: Researchers at Silent Break Security / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Исследователи отправили через систему множество писем, чтобы собрать выходные данные модели из заголовков.</p></a>
<a class="relation-item" href="/studies/AML.CS0009/"><span class="relation-id">AML.CS0009</span><strong>Отравление Tay</strong><span class="relation-meta">Актор: 4chan Users / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Злоумышленники могли взаимодействовать с Tay через сообщения в Twitter.</p></a>
<a class="relation-item" href="/studies/AML.CS0014/"><span class="relation-id">AML.CS0014</span><strong>Сбивание с толку антивирусных нейронных сетей</strong><span class="relation-meta">Актор: Kaspersky ML Research Team / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>На протяжении кейса исследователи использовали доступ к целевому антивирусному продукту на основе ML. Продукт сканирует файлы на системе пользователя, локально извлекает признаки, а затем отправляет их в облачный ML-детектор вредоносного ПО для классификации. Поэтому у исследователей был только доступ к самому детектору в режиме чёрного ящика, но из экстрактора признаков они могли получить ценную информацию для построения атаки.</p></a>
<a class="relation-item" href="/studies/AML.CS0016/"><span class="relation-id">AML.CS0016</span><strong>Выполнение кода в MathGPT через промпт-инъекцию</strong><span class="relation-meta">Актор: Ludwig-Ferdinand Stumpp / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Исследователь смог взаимодействовать с базовой моделью GPT-3 через приложение MathGPT. MathGPT использует GPT-3 для генерации Python-кода, решающего математические задачи, описанные пользовательскими промптами, и показывает пользователю сгенерированный код вместе с решением. Изучение встроенных и пользовательских промптов, а также их результатов привело исследователя к предположению, что приложение напрямую выполняет код, сгенерированный GPT-3.</p></a>
<a class="relation-item" href="/studies/AML.CS0017/"><span class="relation-id">AML.CS0017</span><strong>Обход проверки личности ID.me</strong><span class="relation-meta">Актор: One individual / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Мужчина подал заявки на пособие по безработице в California Employment Development Department, используя поддельные личности, и в процессе взаимодействовал с системой проверки личности ID.me. Система извлекает данные из фотографии документа, проверяет подлинность документа с помощью сочетания ИИ и проприетарных методов, а затем выполняет распознавание лица, сопоставляя фотографию в документе с селфи. &lt;sup&gt;[[7]](https://network.id.me/wp-content/uploads/Document-Verification-Use-Machine-Vision-and-AI-to-Extract-Content-and-Verify-the-Authenticity-1.pdf)&lt;/sup&gt; Мужчина установил, что California Employment Development Department использует сторонний сервис ID.me для проверки личности заявителей. На сайте ID.me описаны шаги проверки личности, включая ввод персональных данных, загрузку водительского удостоверения и отправку селфи.</p></a>
<a class="relation-item" href="/studies/AML.CS0026/"><span class="relation-id">AML.CS0026</span><strong>Перехват финансовой транзакции с использованием M365 Copilot в роли инсайдера</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Во время разработки и выполнения атаки на систему жертвы исследователи Zenity взаимодействовали с Microsoft Copilot for M365.</p></a>
<a class="relation-item" href="/studies/AML.CS0033/"><span class="relation-id">AML.CS0033</span><strong>Обход мобильной KYC-верификации с помощью дипфейк-изображения в реальном времени</strong><span class="relation-meta">Актор: iProov Red Team / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Во время проверки личности приложение финансового сервиса использует распознавание лица и проверку присутствия живого человека для анализа видео с камеры пользователя в реальном времени.</p></a>
<a class="relation-item" href="/studies/AML.CS0034/"><span class="relation-id">AML.CS0034</span><strong>ProKYC: дипфейк-инструмент для атак с мошенническим созданием аккаунтов</strong><span class="relation-meta">Актор: ProKYC, cybercriminal group / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Во время проверки личности приложение финансового сервиса использовало распознавание лица и проверку присутствия живого человека для анализа видео с камеры пользователя в реальном времени.</p></a>
<a class="relation-item" href="/studies/AML.CS0035/"><span class="relation-id">AML.CS0035</span><strong>Эксфильтрация данных из Slack AI через косвенную промпт-инъекцию</strong><span class="relation-meta">Актор: PromptArmor / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Исследователь взаимодействует со Slack AI, отправляя сообщения в публичные каналы Slack.</p></a>
<a class="relation-item" href="/studies/AML.CS0036/"><span class="relation-id">AML.CS0036</span><strong>AIKatz: атака на десктопные LLM-приложения</strong><span class="relation-meta">Актор: Lumia Security / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>После этого злоумышленник получил доступ, необходимый для взаимодействия с backend-сервисом LLM так, будто он является десктопным клиентом. Это давало ему доступ ко всем действиям, которые пользователь может выполнять через десктопное приложение.</p></a>
<a class="relation-item" href="/studies/AML.CS0037/"><span class="relation-id">AML.CS0037</span><strong>Эксфильтрация данных через инструменты ИИ-агента в Copilot Studio</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Далее исследователи повторяют те же шаги для взаимодействия с ИИ-агентом: отправляют агенту вредоносные промпты по электронной почте и получают ответы на нужный им адрес.</p></a>
<a class="relation-item" href="/studies/AML.CS0060/"><span class="relation-id">AML.CS0060</span><strong>Межсайтовый скриптинг (XSS) через манипуляцию промптом в ИИ-чат-боте Lenovo</strong><span class="relation-meta">Актор: Cybernews Research Team / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Исследователи использовали публичный веб-чат с агентом клиентской поддержки Lenovo «Lena».</p></a>
<a class="relation-item" href="/studies/AML.CS0061/"><span class="relation-id">AML.CS0061</span><strong>AI in the Middle: веб-сервисы ИИ как ретрансляторы C2</strong><span class="relation-meta">Актор: Check Point Research / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Предполагая наличие предварительного доступа, исследователи использовали собственный имплант на C++ со встроенным браузером или компонентом WebView для взаимодействия с Grok или Microsoft Copilot через публичный веб-интерфейс, а не через API-ключ.</p></a>
<a class="relation-item" href="/studies/AML.CS0062/"><span class="relation-id">AML.CS0062</span><strong>RCE-уязвимость в Semantic Kernel Search Plugin</strong><span class="relation-meta">Актор: Microsoft Defender Security Research Team / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Исследователи взаимодействовали с агентом на базе Semantic Kernel через его стандартный чат-интерфейс.</p></a>
</div>

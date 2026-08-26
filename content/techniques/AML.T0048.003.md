---
atlas_id: AML.T0048.003
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2023-10-25"
description: Ущерб пользователям может включать разные виды вреда, в том числе финансовый и репутационный, которые направлены на отдельных жертв атаки или ощущаются ими, а не проявляются на уровне организации.
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
    - Enterprise
procedure_count: 16
source_name: User Harm
subtechnique_count: 0
subtechnique_of: AML.T0048
tactics:
    - AML.TA0011
title: Ущерб пользователям
url: /techniques/AML.T0048.003/
---

Ущерб пользователям может включать разные виды вреда, в том числе финансовый и репутационный, которые направлены на отдельных жертв атаки или ощущаются ими, а не проявляются на уровне организации.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0011/"><span class="relation-id">AML.TA0011</span><strong>Воздействие</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0048/"><span class="relation-id">AML.T0048</span><strong>Внешний ущерб</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0048.000/"><span class="relation-id">AML.T0048.000</span><strong>Финансовый ущерб</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0048.001/"><span class="relation-id">AML.T0048.001</span><strong>Репутационный ущерб</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0048.002/"><span class="relation-id">AML.T0048.002</span><strong>Общественный вред</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0048.004/"><span class="relation-id">AML.T0048.004</span><strong>Кража интеллектуальной собственности ИИ</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0020/"><span class="relation-id">AML.CS0020</span><strong>Угрозы косвенной промпт-инъекции: Bing Chat как похититель данных</strong><span class="relation-meta">Актор: Kai Greshake, Saarland University / Тактика: AML.TA0011 Воздействие</span><p>Получив персональные данные пользователя, злоумышленник мог использовать их для дальнейших атак, включая кражу личности или мошенничество.</p></a>
<a class="relation-item" href="/studies/AML.CS0021/"><span class="relation-id">AML.CS0021</span><strong>Эксфильтрация разговоров ChatGPT</strong><span class="relation-meta">Актор: Embrace The Red / Тактика: AML.TA0011 Воздействие</span><p>В результате нарушается конфиденциальность пользователя, и он может стать целью дальнейших атак.</p></a>
<a class="relation-item" href="/studies/AML.CS0022/"><span class="relation-id">AML.CS0022</span><strong>Галлюцинация пакетов ChatGPT</strong><span class="relation-meta">Актор: Vulcan Cyber, Lasso Security / Тактика: AML.TA0011 Воздействие</span><p>Это может привести к различному ущербу для конечного пользователя или организации.</p></a>
<a class="relation-item" href="/studies/AML.CS0024/"><span class="relation-id">AML.CS0024</span><strong>Червь Morris II: атака на основе RAG</strong><span class="relation-meta">Актор: Stav Cohen, Ron Bitton, Ben Nassi / Тактика: AML.TA0011 Воздействие</span><p>Персональные данные пользователей GenAI-почтового ассистента могут утечь к злоумышленникам.</p></a>
<a class="relation-item" href="/studies/AML.CS0029/"><span class="relation-id">AML.CS0029</span><strong>Эксфильтрация разговоров Google Bard</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0011 Воздействие</span><p>Разговор пользователя эксфильтруется, что нарушает его конфиденциальность и может позволить проводить дальнейшие целевые атаки.</p></a>
<a class="relation-item" href="/studies/AML.CS0032/"><span class="relation-id">AML.CS0032</span><strong>Попытка обхода ML-системы обнаружения фишинговых веб-страниц</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0011 Воздействие</span><p>Конечный пользователь может столкнуться с разными видами ущерба, включая финансовый ущерб и нарушение конфиденциальности, в зависимости от того, какие учетные данные украл злоумышленник.</p></a>
<a class="relation-item" href="/studies/AML.CS0036/"><span class="relation-id">AML.CS0036</span><strong>AIKatz: атака на десктопные LLM-приложения</strong><span class="relation-meta">Актор: Lumia Security / Тактика: AML.TA0011 Воздействие</span><p>Злоумышленник мог получить доступ ко всей активности жертвы при работе с LLM, включая прошлые и текущие чаты, а также любые загруженные файлы или содержимое.</p></a>
<a class="relation-item" href="/studies/AML.CS0040/"><span class="relation-id">AML.CS0040</span><strong>Взлом памяти ChatGPT с помощью промпт-инъекции</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0011 Воздействие</span><p>Из-за отравленных воспоминаний ChatGPT жертва может получать ложную информацию, быть введена в заблуждение или подвергаться влиянию.</p></a>
<a class="relation-item" href="/studies/AML.CS0041/"><span class="relation-id">AML.CS0041</span><strong>Бэкдор в файле правил: атака на цепочку поставки ИИ-ассистентов для программирования</strong><span class="relation-meta">Актор: Pillar Security / Тактика: AML.TA0011 Воздействие</span><p>Разработчики-жертвы неосознанно использовали скомпрометированного ИИ-ассистента для программирования, который генерировал код со скрытыми вредоносными элементами, включая бэкдоры, код для эксфильтрации данных, уязвимые конструкции или вредоносные скрипты. Такой код мог попасть в продакшен-приложение и повлиять на пользователей ПО.</p></a>
<a class="relation-item" href="/studies/AML.CS0048/"><span class="relation-id">AML.CS0048</span><strong>Публично доступные интерфейсы управления ClawdBot позволили получить учетные данные и выполнить команды</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0011 Воздействие</span><p>Исследователь мог бы использовать обнаруженные токены приложений для дальнейшего вреда пользователю, включая имперсонацию через отправку сообщений от имени пользователя в любом из подключенных мессенджеров.</p></a>
<a class="relation-item" href="/studies/AML.CS0054/"><span class="relation-id">AML.CS0054</span><strong>Эксфильтрация данных через удаленный отравленный MCP-инструмент</strong><span class="relation-meta">Актор: Invariant Labs / Тактика: AML.TA0011 Воздействие</span><p>Приватные данные пользователя были раскрыты удаленному MCP-серверу.</p></a>
<a class="relation-item" href="/studies/AML.CS0056/"><span class="relation-id">AML.CS0056</span><strong>Кампании по дистилляции моделей, нацеленные на Anthropic Claude</strong><span class="relation-meta">Актор: DeepSeek, Moonshot AI, MiniMax / Тактика: AML.TA0011 Воздействие</span><p>У дистиллированных моделей нет защитных ограничений Claude, из-за чего пользователи могут столкнуться с вредоносными ответами и поведением модели.</p></a>
<a class="relation-item" href="/studies/AML.CS0063/"><span class="relation-id">AML.CS0063</span><strong>Атаки на Gemini с помощью промптов в приглашениях Google Calendar</strong><span class="relation-meta">Актор: SafeBreach Research Team / Тактика: AML.TA0011 Воздействие</span><p>Gemini выдавал жертве контент, способный причинить вред, или многократно показывал рекламные материалы, выбранные злоумышленником.</p></a>
<a class="relation-item" href="/studies/AML.CS0063/"><span class="relation-id">AML.CS0063</span><strong>Атаки на Gemini с помощью промптов в приглашениях Google Calendar</strong><span class="relation-meta">Актор: SafeBreach Research Team / Тактика: AML.TA0011 Воздействие</span><p>Физическая среда жертвы была изменена, что могло создать угрозу её безопасности и приватности, а также причинить имущественный или финансовый ущерб.</p></a>
<a class="relation-item" href="/studies/AML.CS0063/"><span class="relation-id">AML.CS0063</span><strong>Атаки на Gemini с помощью промптов в приглашениях Google Calendar</strong><span class="relation-meta">Актор: SafeBreach Research Team / Тактика: AML.TA0011 Воздействие</span><p>Неавторизованная трансляция видео нарушила приватность жертвы.</p></a>
<a class="relation-item" href="/studies/AML.CS0064/"><span class="relation-id">AML.CS0064</span><strong>Отравленные шаблоны GGUF: атака на цепочку поставок во время инференса</strong><span class="relation-meta">Актор: Pillar Security, Fujitsu Research of Europe / Тактика: AML.TA0011 Воздействие</span><p>The compromised system may expose users to privacy loss, credential theft, misleading information, or attacker-modified software. Malicious content inserted into generated artifacts may continue to affect downstream users after those artifacts are deployed or distributed.</p></a>
</div>

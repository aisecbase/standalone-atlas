---
atlas_id: AML.T0054
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2023-10-25"
description: Злоумышленники могут побуждать большую языковую модель (LLM) игнорировать, обходить или переопределять её поведение в части безопасности/выравнивания и/или гардрейлы, чтобы получать ответы, которые по замыслу модель...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 4
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 9
source_name: LLM Jailbreak
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0007
    - AML.TA0012
title: Джейлбрейк LLM
url: /techniques/AML.T0054/
---

Злоумышленники могут побуждать большую языковую модель (LLM) игнорировать, обходить или переопределять её поведение в части безопасности/выравнивания и/или гардрейлы, чтобы получать ответы, которые по замыслу модель не должна предоставлять. После джейлбрейка злоумышленник может использовать LLM непредусмотренными способами. Джейлбрейка можно добиться с помощью состязательного промптинга либо путём изменения весов модели или механизмов безопасности.

Злоумышленники могут пытаться выполнить джейлбрейк LLM для [уклонения от защиты](/tactics/AML.TA0007), обходя правила и гардрейлы самой LLM, чтобы затем раскрыть информацию (например: [утечка данных из LLM](/techniques/AML.T0057), [выявление системной информации LLM](/techniques/AML.T0069)) или сгенерировать вредоносный контент (например: [генерация вредоносных команд](/techniques/AML.T0102), [целевой фишинг через LLM для социальной инженерии](/techniques/AML.T0052.000)). Они также могут выполнять джейлбрейк модели для [повышения привилегий](/tactics/AML.TA0012), чтобы вызывать инструменты или выполнять действия в собственных целях (например: [вызов инструментов ИИ-агента](/techniques/AML.T0053)), либо использовать агента как канал [командования и управления](/tactics/AML.TA0014) (например: [ИИ-агент](/techniques/AML.T0108)).

Злоумышленники используют различные стратегии для составления джейлбрейк-промптов. Промпты могут быть нацелены на конкретные модели или семейства моделей и итеративно дорабатываться до успешного результата. Поставщики активно обновляют гардрейлы моделей, чтобы повышать устойчивость этих моделей к новым джейлбрейк-промптам. Распространённые стратегии [[jailbreak-guide]] включают, помимо прочего:

- Переопределение инструкций: использовать формулировки, которые пытаются заменить собой ранее заданные ограничения (например, "ignore previous instructions").

- Ролевая игра / смена персоны: предписать LLM принять определённую роль или перейти в режим, допускающий ответы без ограничений (например, "as a security researcher").

- Вымышленный контекст и гипотетические сценарии: предписать LLM включить запрещённый контент в рассказ, сценарий или учебную ситуацию.

- Отделение намерения от содержимого: запросить анализ, примеры, шаблоны или пограничные случаи, которые неявно содержат запрещённый контент.

- Многоходовая эскалация / Crescendo: использовать последовательность промптов, которые начинаются с безобидных запросов, создают доверительный контекст, а затем с каждым новым промптом постепенно пересекают границы политики.

- Ограниченные форматы вывода: предписать LLM выводить результат по строгой схеме или в заданном формате (например, JSON, YAML, код или таблицы).

- Обфускация и преобразование: использовать кодирование, преобразования, перевод или эвфемизмы (например, кодирование base64, "describe it in another language").

- Создание высокоприоритетной цели: представить выполнение запроса как необходимое для основной задачи пользователя (например, "to complete the evaluation," "to follow the spec," "to follow safety guidelines").

- Утвердительная формулировка: добавление в конец промптов утвердительных слов, таких как "sure", может помогать обходить отказы модели и добиваться генерации вредоносного или иного нежелательного контента.[[cybernews]]

Злоумышленники также могут использовать алгоритмические подходы к генерации джейлбрейк-промптов [[jailbreak-zoo]] [[jailbreak-survey]]. Алгоритмическая генерация джейлбрейков позволяет применять автоматизированные методы, которые находят джейлбрейки в больших масштабах. Некоторые подходы автоматизируют ручные стратегии [[autodan]] [[gptfuzzer]] [[crescendo]] [[echo-chamber]], тогда как другие напрямую оптимизируют строку токенов [[universal]], создавая бессмысленный текст. Применимы как подходы к оптимизации в режиме чёрного ящика — для коммерческих моделей, когда у злоумышленника есть только возможность отправлять модели запросы, — так и подходы в режиме белого ящика — для моделей с открытым исходным кодом, когда у злоумышленника есть полный доступ к весам модели.

Злоумышленники также могут напрямую манипулировать весами модели либо изменять или удалять части модели, чтобы создать вариант целевой модели после джейлбрейка или вариант «без цензурных ограничений». Это применимо к моделям с открытым исходным кодом или к случаям, когда злоумышленник получает полный доступ к целевой модели. Подходы включают дообучение для снижения числа отказов [[single-direction]], целевое редактирование модели [[rome]], добавление адаптеров [[lora]] и удаление механизмов безопасности, таких как гардрейлы.

Джейлбрейк-промпты, которые, как известно, работают с различными классами LLM, часто публикуются open-source-сообществом [[dan]]. LLM после джейлбрейка или без цензурных ограничений, обученные или дообученные для работы со снятыми ограничениями, распространяются в публичных реестрах моделей, таких как Hugging Face [[abliteration]].


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0007/"><span class="relation-id">AML.TA0007</span><strong>Уклонение от защиты</strong></a>
<a class="relation-item" href="/tactics/AML.TA0012/"><span class="relation-id">AML.TA0012</span><strong>Повышение привилегий</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0020/"><span class="relation-id">AML.M0020</span><strong>Гардрейлы для генеративного ИИ</strong><p>Гардрейлы могут предотвращать вредоносные входные данные, способные привести к джейлбрейку.</p></a>
<a class="relation-item" href="/mitigations/AML.M0021/"><span class="relation-id">AML.M0021</span><strong>Правила и инструкции для генеративного ИИ</strong><p>Инструкции для модели могут предписывать ей отказываться отвечать на небезопасные входные данные.</p></a>
<a class="relation-item" href="/mitigations/AML.M0022/"><span class="relation-id">AML.M0022</span><strong>Выравнивание модели генеративного ИИ</strong><p>Выравнивание может повысить безопасность, заложенную в параметры модели, снижая её восприимчивость к небезопасным промптам и вероятность формирования небезопасных ответов.</p></a>
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>Красная команда по ИИ</strong><p>Отрабатывайте различные варианты джейлбрейка: выполняемые вручную и автоматизированные, многоходовые, многоязычные, закодированные, преобразованные и мультимодальные. Используйте успешные случаи джейлбрейка для совершенствования гардрейлов, рекомендаций, выравнивания модели и мониторинга, а также включайте их в регрессионные оценки.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0041/"><span class="relation-id">AML.CS0041</span><strong>Бэкдор в файле правил: атака на цепочку поставки ИИ-ассистентов для программирования</strong><span class="relation-meta">Актор: Pillar Security / Тактика: AML.TA0007 Уклонение от защиты</span><p>Промпт использовал техники джейлбрейка, чтобы убедить ИИ-ассистента программирования добавить вредоносный скрипт в сгенерированные HTML-файлы.</p></a>
<a class="relation-item" href="/studies/AML.CS0046/"><span class="relation-id">AML.CS0046</span><strong>Уничтожение данных через косвенную промпт-инъекцию, нацеленную на Claude Computer Use</strong><span class="relation-meta">Актор: HiddenLayer / Тактика: AML.TA0007 Уклонение от защиты</span><p>Промпт сообщал Claude, что он находится в виртуальной среде для тестирования безопасности и что выполнение потенциально опасных команд допустимо. Это позволило обойти гардрейлы Claude, препятствующие выполнению обфусцированных команд.</p></a>
<a class="relation-item" href="/studies/AML.CS0051/"><span class="relation-id">AML.CS0051</span><strong>Использование OpenClaw для командования и управления через промпт-инъекцию</strong><span class="relation-meta">Актор: HiddenLayer / Тактика: AML.TA0007 Уклонение от защиты</span><p>Злоумышленник использовал управляющие последовательности `&lt;think&gt;`, чтобы имитировать внутренние рассуждения и обойти выравнивание модели с требованиями безопасности.</p></a>
<a class="relation-item" href="/studies/AML.CS0052/"><span class="relation-id">AML.CS0052</span><strong>LLMSmith: уязвимости RCE в приложениях с интеграцией LLM</strong><span class="relation-meta">Актор: Researchers at University of Chinese Academy of Sciences, Shandong University, and University of New South Wales / Тактика: AML.TA0007 Уклонение от защиты</span><p>В целевых приложениях, где ИИ-агент отказывался выполнять запрос исследователей, они использовали несложные стратегии джейлбрейка, чтобы обойти гардрейлы LLM.</p></a>
<a class="relation-item" href="/studies/AML.CS0057/"><span class="relation-id">AML.CS0057</span><strong>Storm-2139: обход гардрейлов Azure OpenAI</strong><span class="relation-meta">Актор: Storm-2139 / Тактика: AML.TA0007 Уклонение от защиты</span><p>Storm-2139 намеренно обходила защитные механизмы и фильтры контента Azure OpenAI Service, чтобы генерировать запрещенные результаты. Microsoft сообщила, что злоумышленники итеративно дорабатывали заблокированные промпты, подставляли описания знаменитостей и использовали измененные формулировки или техническую нотацию для обхода фильтров.</p></a>
<a class="relation-item" href="/studies/AML.CS0063/"><span class="relation-id">AML.CS0063</span><strong>Атаки на Gemini с помощью промптов в приглашениях Google Calendar</strong><span class="relation-meta">Актор: SafeBreach Research Team / Тактика: AML.TA0007 Уклонение от защиты</span><p>Вредоносный промпт заставил Gemini следовать ролевому сценарию и указаниям по переопределению инструкций и вместо обычного ответа выдавать выбранное злоумышленником содержимое: токсичный контент или рекламные материалы.</p></a>
<a class="relation-item" href="/studies/AML.CS0066/"><span class="relation-id">AML.CS0066</span><strong>ZombieAgent: атака на ChatGPT с эксфильтрацией данных</strong><span class="relation-meta">Актор: Radware Security Researchers / Тактика: AML.TA0007 Уклонение от защиты</span><p>Исследователи обошли ограничение ChatGPT, запрещавшее динамически формировать или изменять URL, предоставив индексированный словарь статических URL. ChatGPT кодировал собранные данные, выбирая и открывая URL, соответствующий каждому символу и его позиции.</p></a>
<a class="relation-item" href="/studies/AML.CS0067/"><span class="relation-id">AML.CS0067</span><strong>Раскрытие секретов через Claude Code GitHub Action</strong><span class="relation-meta">Актор: Microsoft Defender Security Research Team / Тактика: AML.TA0007 Уклонение от защиты</span><p>Промпт был оформлен как безобидная проверка соблюдения требований и содержал инструкции удалить префикс учётных данных, чтобы обойти поведение Claude при отказе от вывода API-ключа в распознаваемом формате.</p></a>
<a class="relation-item" href="/studies/AML.CS0069/"><span class="relation-id">AML.CS0069</span><strong>Кибершпионская кампания GTG-1002 с использованием Claude Code</strong><span class="relation-meta">Актор: GTG-1002 / Тактика: AML.TA0007 Уклонение от защиты</span><p>Вводящие в заблуждение промпты позволили обойти защитные механизмы Claude Code и побудили его к наступательным действиям, от выполнения которых он по замыслу должен был отказываться.</p></a>
</div>


## Источники

- [Uncensor any LLM with abliteration](https://huggingface.co/blog/mlabonne/abliteration)
- [AutoDAN: Generating Stealthy Jailbreak Prompts on Aligned Large Language Models](https://arxiv.org/abs/2310.04451)
- [Great, Now Write an Article About That: The Crescendo Multi-Turn LLM Jailbreak Attack](https://arxiv.org/abs/2404.01833)
- [ChatGPT DAN](https://github.com/0xk1h0/ChatGPT_DAN)
- [The Echo Chamber Multi-Turn LLM Jailbreak](https://arxiv.org/abs/2601.05742)
- [GPTFUZZER: Red Teaming Large Language Models with Auto-Generated Jailbreak Prompts](https://arxiv.org/abs/2309.10253)
- [Jailbreaking LLMs: A Comprehensive Guide (With Examples)](https://www.promptfoo.dev/blog/how-to-jailbreak-llms/)
- [Jailbreak Attacks and Defenses Against Large Language Models: A Survey](https://arxiv.org/abs/2407.04295)
- [JailbreakZoo: Survey, Landscapes, and Horizons in Jailbreaking Large Language and Vision-Language Models](https://arxiv.org/abs/2407.01599)
- [LoRA Fine-tuning Efficiently Undoes Safety Training in Llama 2-Chat 70B](https://arxiv.org/abs/2310.20624)
- [Locating and Editing Factual Associations in GPT](https://arxiv.org/abs/2202.05262)
- [Refusal in Language Models Is Mediated by a Single Direction](https://arxiv.org/abs/2406.11717)
- [Universal and Transferable Adversarial Attacks on Aligned Language Models](https://arxiv.org/abs/2307.15043)
- [GitHub Copilot Jailbreak Vulnerability Let Attackers Train Malicious Models](https://cybersecuritynews.com/github-copilot-jailbreak-vulnerability)

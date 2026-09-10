---
atlas_id: AML.T0017
atlas_type: technique
attack_ref_id: T1587
attack_ref_url: https://attack.mitre.org/techniques/T1587/
created_date: "2023-10-25"
description: Злоумышленники могут разрабатывать собственные средства для проведения своих операций. Этот процесс охватывает определение требований, создание или адаптацию решений, проверку или упаковку средств, а также их...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-08-31"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
    - Enterprise
procedure_count: 10
source_name: Develop Capabilities
subtechnique_count: 3
subtechnique_of: ""
tactics:
    - AML.TA0003
title: Разработка средств для атаки
url: /techniques/AML.T0017/
---

Злоумышленники могут разрабатывать собственные средства для проведения своих операций. Этот процесс охватывает определение требований, создание или адаптацию решений, проверку или упаковку средств, а также их подготовку к развёртыванию. Средства, используемые для проведения атак на системы с поддержкой ИИ, не обязательно сами основаны на ИИ. Злоумышленники также могут задействовать автономных ИИ-агентов для итеративной разработки средств, в том числе методов эксплуатации известных или ранее неизвестных уязвимостей ПО.

К примерам относятся создание состязательных атак на ИИ, разработка веб-сайтов с вредоносными инструкциями для ИИ-агентов, создание вредоносных ИИ-артефактов, разработка вредоносного ПО и реализация эксплойтов для ПО.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0017.000/"><span class="relation-id">AML.T0017.000</span><strong>Состязательные атаки на ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0017.001/"><span class="relation-id">AML.T0017.001</span><strong>Автономная разработка эксплойтов</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0017.002/"><span class="relation-id">AML.T0017.002</span><strong>Инструменты ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0018/"><span class="relation-id">AML.CS0018</span><strong>Выполнение произвольного кода через Google Colab</strong><span class="relation-meta">Актор: Tony Piazza / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленник создает Jupyter Notebook, содержащий обфусцированный вредоносный код.</p></a>
<a class="relation-item" href="/studies/AML.CS0020/"><span class="relation-id">AML.CS0020</span><strong>Угрозы косвенной промпт-инъекции: Bing Chat как похититель данных</strong><span class="relation-meta">Актор: Kai Greshake, Saarland University / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленник создал сайт с вредоносными системными инструкциями для LLM, чтобы повлиять на поведение модели. Эти инструкции считываются моделью, когда пользователь предоставляет Bing Chat доступ к открытой странице.</p></a>
<a class="relation-item" href="/studies/AML.CS0029/"><span class="relation-id">AML.CS0029</span><strong>Эксфильтрация разговоров Google Bard</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователь написал Google Apps Script, который записывает все параметры запроса в Google Doc.</p></a>
<a class="relation-item" href="/studies/AML.CS0043/"><span class="relation-id">AML.CS0043</span><strong>Прототип вредоносного ПО со встроенной промпт-инъекцией</strong><span class="relation-meta">Актор: Unknown Threat Actor / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленник встроил промпт-инъекцию в образец вредоносного ПО, названный Skynet.</p></a>
<a class="relation-item" href="/studies/AML.CS0049/"><span class="relation-id">AML.CS0049</span><strong>Компрометация цепочки поставки через отравленный навык ClawdBot</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователь создал простой веб-сервер для записи запросов.</p></a>
<a class="relation-item" href="/studies/AML.CS0050/"><span class="relation-id">AML.CS0050</span><strong>Удаленное выполнение кода (RCE) в OpenClaw в один клик</strong><span class="relation-meta">Актор: DepthFirst / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователь разработал JavaScript-скрипт для RCE в один клик.</p></a>
<a class="relation-item" href="/studies/AML.CS0052/"><span class="relation-id">AML.CS0052</span><strong>LLMSmith: уязвимости RCE в приложениях с интеграцией LLM</strong><span class="relation-meta">Актор: Researchers at University of Chinese Academy of Sciences, Shandong University, and University of New South Wales / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи провели статический анализ API целевых LLM-фреймворков, чтобы выявить функции, которые выполняют код из пользовательского ввода или ответа LLM и поэтому уязвимы к RCE.</p></a>
<a class="relation-item" href="/studies/AML.CS0053/"><span class="relation-id">AML.CS0053</span><strong>Эксфильтрация писем через отравленный MCP-сервер Postmark</strong><span class="relation-meta">Актор: Unknown Bad Actor / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленник изменил легитимный MCP-сервер Postmark так, чтобы его адрес электронной почты добавлялся в поле скрытой копии (BCC) всех писем, отправляемых инструментом.</p></a>
<a class="relation-item" href="/studies/AML.CS0055/"><span class="relation-id">AML.CS0055</span><strong>AI ClickFix: захват управления computer-use-агентами с помощью ClickFix</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователь использовал ChatGPT, чтобы с небольшими ручными правками сгенерировать вредоносный сайт. Сайт был рассчитан на то, чтобы побудить computer-use-агентов взаимодействовать с определенными элементами и в итоге выполнить код исследователя на машине жертвы. Исследователь также написал скрипт, способный запускать приложение на машине жертвы.</p></a>
<a class="relation-item" href="/studies/AML.CS0057/"><span class="relation-id">AML.CS0057</span><strong>Storm-2139: обход гардрейлов Azure OpenAI</strong><span class="relation-meta">Актор: Storm-2139 / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Создатели Storm-2139 разработали инструмент под названием de3u, чтобы упростить неавторизованное использование сервисов генеративного ИИ и обход защитных механизмов.</p></a>
</div>

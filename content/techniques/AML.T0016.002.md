---
atlas_id: AML.T0016.002
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-03-12"
description: Злоумышленники могут искать и получать модели или инструменты генеративного ИИ, например большие языковые модели (LLM), чтобы использовать их на различных этапах своих операций. Генеративный ИИ может использоваться во...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 3
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 4
source_name: Generative AI
subtechnique_count: 0
subtechnique_of: AML.T0016
tactics:
    - AML.TA0003
title: Генеративный ИИ
url: /techniques/AML.T0016.002/
---

Злоумышленники могут искать и получать модели или инструменты генеративного ИИ, например большие языковые модели (LLM), чтобы использовать их на различных этапах своих операций. Генеративный ИИ может использоваться во вредоносных целях, например для генерации вредоносного ПО, [создания дипфейков](/techniques/AML.T0088), [генерации вредоносных команд](/techniques/AML.T0102), [подготовки содержимого для извлечения](/techniques/AML.T0066) или создания контента для [фишинга](/techniques/AML.T0052).

Злоумышленники могут получать модели с открытым исходным кодом и локально развёртывать их с помощью таких фреймворков, как [Ollama](https://ollama.com/) или [vLLM]( https://docs.vllm.ai/en/latest/). Они могут размещать эти модели в облачной инфраструктуре. Кроме того, они могут пользоваться услугами поставщиков ИИ-сервисов, таких как Hugging Face.

Им может потребоваться выполнить джейлбрейк модели (см. [джейлбрейк LLM](/techniques/AML.T0054)), чтобы обойти ограничения на типы генерируемых ею ответов. Им также может потребоваться нарушить установленные разработчиком модели условия использования сервиса.

Модели генеративного ИИ также могут быть «без цензурных ограничений»: они спроектированы для генерации контента без каких-либо ограничений, таких как гардрейлы или контентные фильтры. Генеративный ИИ без цензурных ограничений предоставляет киберпреступникам широкие возможности для злоупотреблений [[blog]] [[gbhackers]]. Модели могут быть дообучены, чтобы устранить выравнивание модели и гардрейлы [[erichartford]], либо подвергнуты целенаправленным манипуляциям для обхода отказа модели выполнять запросы [[arxiv]]; в результате появляются варианты модели без цензурных ограничений. Модели без цензурных ограничений могут создаваться для наступательных и защитных задач в области кибербезопасности [[taico]], а злоумышленник может использовать их возможности во вредоносных целях. Существуют также модели, специально разработанные и рекламируемые для вредоносного использования [[gbhackers-1]].


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0016/"><span class="relation-id">AML.T0016</span><strong>Получение средств для атаки</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0016.000/"><span class="relation-id">AML.T0016.000</span><strong>Готовые реализации состязательных атак на ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.001/"><span class="relation-id">AML.T0016.001</span><strong>Программные инструменты</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.003/"><span class="relation-id">AML.T0016.003</span><strong>Exploits</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.004/"><span class="relation-id">AML.T0016.004</span><strong>AI Agent Tools</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0001/"><span class="relation-id">AML.M0001</span><strong>Ограничение публикации артефактов модели</strong><p>Ограничивайте публикацию в открытом доступе артефактов моделей, используемых в продакшене: злоумышленники могут получить такие артефакты и адаптировать их для использования в своих операциях.</p></a>
<a class="relation-item" href="/mitigations/AML.M0020/"><span class="relation-id">AML.M0020</span><strong>Гардрейлы для генеративного ИИ</strong><p>Применяйте гардрейлы, чтобы блокировать вредоносное использование моделей или сервисов генеративного ИИ, а также попытки их джейлбрейка.</p></a>
<a class="relation-item" href="/mitigations/AML.M0022/"><span class="relation-id">AML.M0022</span><strong>Выравнивание модели генеративного ИИ</strong><p>Выравнивайте генеративные модели, чтобы они противостояли вредоносным запросам злоумышленника и попыткам устранить предусмотренное безопасное поведение.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0033/"><span class="relation-id">AML.CS0033</span><strong>Обход мобильной KYC-верификации с помощью дипфейк-изображения в реальном времени</strong><span class="relation-meta">Актор: iProov Red Team / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи получили [Faceswap](https://swapface.org) — настольное приложение, способное подменять лица на видео в реальном времени.</p></a>
<a class="relation-item" href="/studies/AML.CS0034/"><span class="relation-id">AML.CS0034</span><strong>ProKYC: дипфейк-инструмент для атак с мошенническим созданием аккаунтов</strong><span class="relation-meta">Актор: ProKYC, cybercriminal group / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленник оплатил инструмент ProKYC, создал поддельное удостоверение личности, сгенерировал дипфейк-видео с селфи и заменил видеопоток с камеры этим дипфейк-видео.</p></a>
<a class="relation-item" href="/studies/AML.CS0055/"><span class="relation-id">AML.CS0055</span><strong>AI ClickFix: захват управления computer-use-агентами с помощью ClickFix</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователь получил доступ к ChatGPT.</p></a>
<a class="relation-item" href="/studies/AML.CS0070/"><span class="relation-id">AML.CS0070</span><strong>Threat Actor Uses a DeepSeek-Powered Hermes Agent in Langflow and n8n Exploitation Attempts</strong><span class="relation-meta">Актор: Chinese-speaking threat actor using the aliases knaithe and KnYuan / Тактика: AML.TA0003 Подготовка ресурсов</span><p>The actor obtained access to several generative-AI models and services while evaluating an operational toolset. DeepSeek was selected as the primary reasoning engine for the autonomous attack activity.</p></a>
</div>


## Источники

- [\[2406.11717\] Refusal in Language Models Is Mediated by a Single Direction](https://arxiv.org/abs/2406.11717/)
- [Cybercriminal abuse of large language models](https://blog.talosintelligence.com/cybercriminal-abuse-of-large-language-models/)
- [erichartford](https://erichartford.com/uncensored-models)
- [Cybercriminals Exploit LLM Models to Enhance Hacking Activities](https://gbhackers.com/cybercriminals-exploit-llm-models/)
- [BlackHat AI Tool WormGPT Enhanced with Grok and Mixtral](https://gbhackers.com/wormgpt-enhanced-with-grok-and-mixtral/)
- [TAICO | WhiteRabbitNeo: An Uncensored, Open Source AI Model for Red & Blue Team Cybersecurity](https://taico.ca/posts/whiterabbitneo/)

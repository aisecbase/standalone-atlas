---
atlas_id: AML.T0010.001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут нацеливаться на программные пакеты, которые широко используются в системах с поддержкой ИИ или являются частью жизненного цикла AI DevOps. Это может включать фреймворки глубокого обучения,...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 3
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 5
source_name: AI Software
subtechnique_count: 0
subtechnique_of: AML.T0010
tactics:
    - AML.TA0004
title: ПО для ИИ
url: /techniques/AML.T0010.001/
---

Злоумышленники могут нацеливаться на программные пакеты, которые широко используются в системах с поддержкой ИИ или являются частью жизненного цикла AI DevOps. Это может включать фреймворки глубокого обучения, используемые для создания ИИ-моделей (например, PyTorch, TensorFlow, Jax), фреймворки интеграции генеративного ИИ (например, LangChain, LangFlow), движки инференса и инструменты AI DevOps. Они также могут атаковать цепочки зависимостей любых из этих программных пакетов [[pytorch]]. Кроме того, злоумышленники могут нацеливаться на отдельные компоненты, используемые ПО для ИИ, например конфигурационные файлы [[pillar]] или примеры использования ИИ-пакетов, которые могут распространяться в Jupyter-ноутбуках [[medium]].

Злоумышленники могут компрометировать легитимные пакеты [[aws]] или публиковать вредоносное ПО под именами, имитирующими легитимные пакеты [[pytorch]]. Они могут выбирать в качестве целей имена пакетов, галлюцинированные большими языковыми моделями [[trendmicro]] (см. [публикацию галлюцинированных сущностей](/techniques/AML.T0060)). Они также могут выполнить [подмену компонента после одобрения в цепочке поставок ИИ](/techniques/AML.T0109): сначала опубликовать легитимный пакет, а затем выпустить вредоносную версию, когда пакет наберет критическую массу пользователей.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0004/"><span class="relation-id">AML.TA0004</span><strong>Первичный доступ</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0010/"><span class="relation-id">AML.T0010</span><strong>Компрометация цепочки поставок ИИ</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0010.000/"><span class="relation-id">AML.T0010.000</span><strong>Аппаратное обеспечение</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.002/"><span class="relation-id">AML.T0010.002</span><strong>Данные</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.003/"><span class="relation-id">AML.T0010.003</span><strong>Модель</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.004/"><span class="relation-id">AML.T0010.004</span><strong>Реестр контейнеров</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.005/"><span class="relation-id">AML.T0010.005</span><strong>Инструмент ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0006/"><span class="relation-id">AML.M0006</span><strong>Использование ансамблевых методов</strong><p>Использование нескольких разных моделей обеспечивает минимальную потерю производительности, если уязвимость обнаружена в инструменте для одной модели или семейства моделей.</p></a>
<a class="relation-item" href="/mitigations/AML.M0013/"><span class="relation-id">AML.M0013</span><strong>Подписание кода</strong><p>Требуйте корректной подписи драйверов и ML-фреймворков.</p></a>
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>Красная команда по ИИ</strong><p>Introduce controlled untrusted AI packages, libraries, plugins, or software components. Verify dependency controls, scanning, signing, approval, isolation, and safe installation.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0015/"><span class="relation-id">AML.CS0015</span><strong>Компрометация цепочки зависимостей PyTorch</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0004 Первичный доступ</span><p>Вредоносный пакет зависимости с именем `torchtriton` был загружен в репозиторий PyPI с тем же именем, что и пакет, поставлявшийся со сборкой PyTorch-nightly. Этот вредоносный пакет содержал дополнительный код, отправлявший чувствительные данные с машины. Вредоносный `torchtriton` устанавливался вместо легитимного пакета, потому что PyPI имел приоритет над другими источниками. Подробнее см. [этот GitHub issue](https://github.com/pypa/pip/issues/8606).</p></a>
<a class="relation-item" href="/studies/AML.CS0018/"><span class="relation-id">AML.CS0018</span><strong>Выполнение произвольного кода через Google Colab</strong><span class="relation-meta">Актор: Tony Piazza / Тактика: AML.TA0004 Первичный доступ</span><p>Jupyter Notebook часто используются для исследований и экспериментов в области ML и data science и содержат исполняемые фрагменты Python-кода, а также типовую функциональность командной строки Unix. Пользователи могут столкнуться со скомпрометированным notebook на публичных сайтах или получить его напрямую по ссылке.</p></a>
<a class="relation-item" href="/studies/AML.CS0022/"><span class="relation-id">AML.CS0022</span><strong>Галлюцинация пакетов ChatGPT</strong><span class="relation-meta">Актор: Vulcan Cyber, Lasso Security / Тактика: AML.TA0004 Первичный доступ</span><p>Пользователь ChatGPT или другой LLM может задать похожий вопрос, получить то же галлюцинированное имя пакета и скачать вредоносный пакет. Исследователи показали, что несколько LLM могут выдавать одни и те же галлюцинации, и зафиксировали более 30 000 скачиваний пакета `huggingface-cli`.</p></a>
<a class="relation-item" href="/studies/AML.CS0041/"><span class="relation-id">AML.CS0041</span><strong>Бэкдор в файле правил: атака на цепочку поставки ИИ-ассистентов для программирования</strong><span class="relation-meta">Актор: Pillar Security / Тактика: AML.TA0004 Первичный доступ</span><p>Исследователи могли бы загрузить вредоносный файл правил в сообщества open source-разработчиков, где конфигурации ИИ-ассистентов программирования распространяются с минимальной проверкой безопасности, например на GitHub и cursor.directory. После включения в репозиторий проекта он может сохраняться при форках проекта и распространении шаблонов, создавая долгосрочную компрометацию цепочки поставки ПО на основе ИИ во многих организациях.</p></a>
<a class="relation-item" href="/studies/AML.CS0047/"><span class="relation-id">AML.CS0047</span><strong>Код для развертывания деструктивного ИИ-агента обнаружен в расширении Amazon Q для VS Code</strong><span class="relation-meta">Актор: lkmanka58 (GitHub user) / Тактика: AML.TA0004 Первичный доступ</span><p>`lkmanka58` использовал GitHub-токен, чтобы добавить вредоносный код в GitHub-репозиторий расширения Amazon Q для VS Code. Коммит автоматически попал в релиз `v1.84.0`.</p></a>
</div>


## Источники

- [Security Update for Amazon Q Developer Extension for Visual Studio Code (Version #1.84)](https://aws.amazon.com/security/security-bulletins/AWS-2025-015/)
- [Careful Who You Colab With: abusing google colaboratory](https://medium.com/mlearning-ai/careful-who-you-colab-with-fa8001f933e7)
- [New Vulnerability in GitHub Copilot and Cursor: How Hackers Can Weaponize Code Agents](https://www.pillar.security/blog/new-vulnerability-in-github-copilot-and-cursor-how-hackers-can-weaponize-code-agents)
- [Compromised PyTorch-nightly dependency chain between December 25th and December 30th, 2022.](https://pytorch.org/blog/compromised-nightly-dependency/)
- [Slopsquatting: When AI Agents Hallucinate Malicious Packages](https://www.trendmicro.com/vinfo/us/security/news/cybercrime-and-digital-threats/slopsquatting-when-ai-agents-hallucinate-malicious-packages)

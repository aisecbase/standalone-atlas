---
atlas_id: AML.T0095
atlas_type: technique
attack_ref_id: T1593
attack_ref_url: https://attack.mitre.org/techniques/T1593/
created_date: "2025-11-05"
description: Злоумышленники могут искать на публичных сайтах и/или доменах сведения о жертвах, которые могут использоваться при выборе цели или подготовке атаки. Информация о жертвах может быть доступна на разных онлайн-площадках,...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 3
source_name: Search Open Websites/Domains
subtechnique_count: 1
subtechnique_of: ""
tactics:
    - AML.TA0002
title: Поиск на открытых сайтах и доменах
url: /techniques/AML.T0095/
---

Злоумышленники могут искать на публичных сайтах и/или доменах сведения о жертвах, которые могут использоваться при выборе цели или подготовке атаки. Информация о жертвах может быть доступна на разных онлайн-площадках, таких как социальные сети, новостные сайты или домены, принадлежащие жертве.

Злоумышленники могут находить нужную им информацию через поисковые системы. Они могут использовать точные поисковые запросы, чтобы выявить программные платформы или сервисы, используемые жертвой, и затем использовать эти сведения при выборе цели или подготовке атаки. За этим может последовать [эксплуатация публичного приложения](/techniques/AML.T0049) или [внедрение промпта через публичное приложение](/techniques/AML.T0093).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0002/"><span class="relation-id">AML.TA0002</span><strong>Разведка</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0095.000/"><span class="relation-id">AML.T0095.000</span><strong>Репозитории кода</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0000/"><span class="relation-id">AML.M0000</span><strong>Ограничение публичного раскрытия информации</strong><p>Ограничивайте размещение на веб-сайтах и доменах общедоступной технической и организационной информации, раскрывающей ИИ-стек, сервисы, сведения о сотрудниках или другие данные, полезные для выбора целей.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0039/"><span class="relation-id">AML.CS0039</span><strong>Living Off AI: промпт-инъекция через Jira Service Management</strong><span class="relation-meta">Актор: Cato CTRL / Тактика: AML.TA0002 Разведка</span><p>Исследователи использовали поисковый запрос `site:atlassian.net/servicedesk inurl:portal`, чтобы выявить организации, использующие сервисные порталы Atlassian как потенциальные цели.</p></a>
<a class="relation-item" href="/studies/AML.CS0061/"><span class="relation-id">AML.CS0061</span><strong>AI in the Middle: веб-сервисы ИИ как ретрансляторы C2</strong><span class="relation-meta">Актор: Check Point Research / Тактика: AML.TA0002 Разведка</span><p>Исследователи оценили публичные ИИ-ассистенты с анонимным или неаутентифицированным веб-просмотром и поведением получения URL, чтобы выявить сервисы, способные запрашивать произвольные URL, подконтрольные злоумышленнику, без API-учетных данных. Исследователи обнаружили, что Grok и Microsoft Copilot соответствуют этим условиям.</p></a>
<a class="relation-item" href="/studies/AML.CS0065/"><span class="relation-id">AML.CS0065</span><strong>Атака на цепочку поставок через повторное использование пространства имён модели</strong><span class="relation-meta">Актор: Unit 42 Researchers / Тактика: AML.TA0002 Разведка</span><p>Специалисты Unit 42 изучили публичные каталоги моделей, использующие Hugging Face, страницы авторов Hugging Face и репозитории проектов с открытым исходным кодом, чтобы выявить ссылки на модели, у которых исходное пространство имён автора было удалено и стало доступно для регистрации. Они выявили сохранявшиеся без обновления ссылки в облачных каталогах, исходном коде, документации, аргументах по умолчанию и ноутбуках с примерами.</p></a>
</div>

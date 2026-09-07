---
atlas_id: AML.T0036
atlas_type: technique
attack_ref_id: T1213
attack_ref_url: https://attack.mitre.org/techniques/T1213/
created_date: "2022-01-24"
description: Злоумышленники могут использовать информационные репозитории для поиска ценных сведений. Информационные репозитории — это инструменты для хранения информации, обычно предназначенные для совместной работы или обмена...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 4
source_name: Data from Information Repositories
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0009
title: Данные из информационных репозиториев
url: /techniques/AML.T0036/
---

Злоумышленники могут использовать информационные репозитории для поиска ценных сведений. Информационные репозитории — это инструменты для хранения информации, обычно предназначенные для совместной работы или обмена данными между пользователями. Они могут содержать широкий спектр данных, которые помогают злоумышленникам достигать дальнейших целей или дают прямой доступ к целевой информации.

Состав информации в репозитории зависит от конкретного экземпляра или среды. К распространенным информационным репозиториям относятся SharePoint, Confluence и корпоративные базы данных, например SQL Server.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0009/"><span class="relation-id">AML.TA0009</span><strong>Сбор материалов</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0006/"><span class="relation-id">AML.CS0006</span><strong>Ошибочная конфигурация Clearview AI</strong><span class="relation-meta">Актор: Researchers at spiderSilk / Тактика: AML.TA0009 Сбор материалов</span><p>Приватный репозиторий кода содержал учетные данные, которые использовались для доступа к облачным хранилищам AWS S3. Это привело к обнаружению ресурсов инструмента распознавания лиц, включая:</p><ul><li>выпущенные настольные и мобильные приложения</li><li>предварительные версии приложений с новыми возможностями</li><li>токены доступа Slack</li><li>необработанные видео и другие данные</li></ul></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0009 Сбор материалов</span><p>Агенты прочитали данные из внутренней рабочей базы данных MongoDB, скачали четыре приватных репозитория исходного кода и получили доступ к пяти клиентским наборам данных, связанным с материалами ExploitGym или CyberGym.</p></a>
<a class="relation-item" href="/studies/AML.CS0069/"><span class="relation-id">AML.CS0069</span><strong>Кибершпионская кампания GTG-1002 с использованием Claude Code</strong><span class="relation-meta">Актор: GTG-1002 / Тактика: AML.TA0009 Сбор материалов</span><p>Аутентифицировавшись с помощью собранных учётных данных, джейлбрейкнутый агент Claude от GTG-1002 обращался с запросами к внутренним базам данных и системам, чтобы получить проприетарную информацию, конфигурации систем и чувствительные эксплуатационные данные.</p></a>
<a class="relation-item" href="/studies/AML.CS0071/"><span class="relation-id">AML.CS0071</span><strong>Multi-Agent Framework Compromises Taiwanese Government Systems</strong><span class="relation-meta">Актор: Unknown Chinese-language actor / Тактика: AML.TA0009 Сбор материалов</span><p>The framework retrieved employee names, departments, identifiers, and SSO account information from an exposed user-database API without authentication.</p></a>
</div>

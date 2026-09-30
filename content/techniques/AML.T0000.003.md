---
atlas_id: AML.T0000.003
atlas_type: technique
attack_ref_id: T1596.005
attack_ref_url: https://attack.mitre.org/techniques/T1596/005/
created_date: "2026-09-15"
description: Злоумышленники могут искать сведения в общедоступных сервисах сканирования интернета, чтобы выявить доступную из интернета инфраструктуру ИИ жертвы. Такие сервисы, как Shodan и Censys, непрерывно сканируют интернет и...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 1
modified_date: "2026-09-15"
platforms:
    - Enterprise
procedure_count: 2
source_name: Scan Databases
subtechnique_count: 0
subtechnique_of: AML.T0000
tactics:
    - AML.TA0002
title: Базы данных с результатами сканирований
url: /techniques/AML.T0000.003/
---

Злоумышленники могут искать сведения в общедоступных сервисах сканирования интернета, чтобы выявить доступную из интернета инфраструктуру ИИ жертвы. Такие сервисы, как Shodan и Censys, непрерывно сканируют интернет и публикуют активные IP-адреса, имена хостов, открытые порты и баннеры сервисов. Злоумышленники могут выполнять поиск по этим данным без прямого взаимодействия с целью.

Собранные таким образом сведения также могут помочь выявить потенциальные цели для последующего применения техники [Активное сканирование](/techniques/AML.T0006), чтобы подтвердить, что сервисы по-прежнему доступны, и проверить наличие ошибок конфигурации или несанкционированных конечных точек. Эти сведения также могут использоваться при подготовке последующих попыток первоначального доступа, например с помощью техники [Эксплуатация приложения, доступного из интернета](/techniques/AML.T0049). В отличие от техники [Активное сканирование](/techniques/AML.T0006), эта техника использует результаты сканирования, полученные третьими сторонами, и не предполагает прямого взаимодействия с системой жертвы.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0002/"><span class="relation-id">AML.TA0002</span><strong>Разведка</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0000/"><span class="relation-id">AML.T0000</span><strong>Поиск в открытых технических базах данных</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0000.000/"><span class="relation-id">AML.T0000.000</span><strong>Журналы и материалы конференций</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0000.001/"><span class="relation-id">AML.T0000.001</span><strong>Репозитории препринтов</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0000.002/"><span class="relation-id">AML.T0000.002</span><strong>Технические блоги</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0039/"><span class="relation-id">AML.M0039</span><strong>AI Honeypots</strong><p>Decoy assets are catalogued by the same internet scan databases (e.g., Shodan, Censys) that adversaries query to locate exposed AI infrastructure. Adversarial engagement that originates from those platforms can grant defenders visibility into those discovery channels.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0048/"><span class="relation-id">AML.CS0048</span><strong>Публично доступные интерфейсы управления ClawdBot позволили получить учётные данные и выполнить команды</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0002 Разведка</span><p>The researcher performed targeting by searching for the title tag of ClawdBot&#39;s web-based control interface, &#34;Clawdbot Control&#34; on Shodan, identifying hundreds of ClawdBot control interfaces exposed on the public internet.</p></a>
<a class="relation-item" href="/studies/AML.CS0070/"><span class="relation-id">AML.CS0070</span><strong>Злоумышленник использовал Hermes Agent на базе DeepSeek при попытках эксплуатации Langflow и n8n</strong><span class="relation-meta">Актор: Chinese-speaking threat actor using the aliases knaithe and KnYuan / Тактика: AML.TA0002 Разведка</span><p>DeepSeek queried FOFA and obtained records for 84 exposed Langflow instances. These were exposure records, not confirmed vulnerable targets.</p></a>
</div>

---
atlas_id: AML.T0097
atlas_type: technique
attack_ref_id: T1497
attack_ref_url: https://attack.mitre.org/techniques/T1497/
created_date: "2025-11-25"
description: Злоумышленники могут применять различные способы, чтобы распознавать и обходить среды виртуализации и анализа. В том числе они могут изменять поведение вредоносного ПО по результатам проверок на наличие артефактов,...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 1
source_name: Virtualization/Sandbox Evasion
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0007
title: Обход виртуализации и песочниц
url: /techniques/AML.T0097/
---

Злоумышленники могут применять различные способы, чтобы распознавать и обходить среды виртуализации и анализа. В том числе они могут изменять поведение вредоносного ПО по результатам проверок на наличие артефактов, характерных для среды виртуальной машины (VME) или песочницы. Если злоумышленник обнаруживает VME, он может изменить вредоносное ПО так, чтобы оно прекратило работу в системе жертвы или скрыло основные функции импланта. Также злоумышленники могут искать артефакты VME перед загрузкой вторичных или дополнительных полезных нагрузок.

Для обхода виртуализации и песочниц злоумышленники могут использовать разные методы: проверять наличие средств мониторинга безопасности, например Sysinternals или Wireshark, а также других системных артефактов, связанных с анализом или виртуализацией. К таким артефактам относятся ключи реестра, например подстроки Vmware, VBOX или QEMU, переменные окружения, например подстроки VBOX, VMWARE или PARALLELS, MAC-адреса сетевых адаптеров, например префиксы 00-05-69 (VMWare) или 08-00-27 (VirtualBox), и запущенные процессы, например vmware.exe, vboxservice.exe или qemu-ga.exe [[research]].


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0007/"><span class="relation-id">AML.TA0007</span><strong>Уклонение от защиты</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0043/"><span class="relation-id">AML.CS0043</span><strong>Прототип вредоносного ПО со встроенной промпт-инъекцией</strong><span class="relation-meta">Актор: Unknown Threat Actor / Тактика: AML.TA0007 Уклонение от защиты</span><p>Вредоносное ПО Skynet пытается применять различные способы обхода песочниц.</p></a>
</div>


## Источники

- [New Malware Embeds Prompt Injection to Evade AI Detection - Check Point Research](https://research.checkpoint.com/2025/ai-evasion-prompt-injection/)

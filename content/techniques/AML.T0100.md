---
atlas_id: AML.T0100
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-11-25"
description: 'Злоумышленники могут создавать обманчивое веб-содержимое, рассчитанное на то, чтобы приманить computer-use-агентов или ИИ-веб-браузеры к выполнению непреднамеренных действий: нажатию кнопок, копированию кода или...'
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Agentic AI
procedure_count: 1
source_name: AI Agent Clickbait
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0005
title: Кликбейт для ИИ-агента
url: /techniques/AML.T0100/
---

Злоумышленники могут создавать обманчивое веб-содержимое, рассчитанное на то, чтобы приманить computer-use-агентов или ИИ-веб-браузеры к выполнению непреднамеренных действий: нажатию кнопок, копированию кода или переходу на определенные веб-страницы. Такие атаки эксплуатируют то, как агент интерпретирует содержимое интерфейса, визуальные подсказки или встроенные в сайт формулировки, похожие на промпты. При успешной атаке агент может непреднамеренно скопировать и выполнить вредоносный код в операционной системе пользователя.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0005/"><span class="relation-id">AML.TA0005</span><strong>Выполнение</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0020/"><span class="relation-id">AML.M0020</span><strong>Гардрейлы для генеративного ИИ</strong><p>Применяйте гардрейлы к недоверенному веб-содержимому и предлагаемым действиям по управлению компьютером.</p></a>
<a class="relation-item" href="/mitigations/AML.M0021/"><span class="relation-id">AML.M0021</span><strong>Правила и инструкции для генеративного ИИ</strong><p>Предписывайте ИИ-агентам считать содержащиеся на веб-страницах инструкции недоверенными и запрашивать одобрение на выполнение действий, способных повлечь существенные последствия.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0055/"><span class="relation-id">AML.CS0055</span><strong>AI ClickFix: захват управления computer-use-агентами с помощью ClickFix</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0005 Выполнение</span><p>Claude Computer-Use Agent жертвы был обманом вовлечен во взаимодействие с вредоносным сайтом текстом:</p></a>
</div>

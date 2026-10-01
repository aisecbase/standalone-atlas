---
atlas_id: AML.M0039
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Technical - AI
    - Technical - Cyber
created_date: "2026-09-15"
description: Deploy monitored decoy AI services, agents, credentials, or other resources that resemble legitimate, active assets in the external or internal environment, but serve no real business or operational functions. Because...
generated: true
generated_by: atlasgen
ml_lifecycle:
    - Deployment
    - Monitoring and Maintenance
modified_date: "2026-09-15"
source_name: AI Honeypots
technique_count: 15
title: AI Honeypots
url: /mitigations/AML.M0039/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Deploy monitored decoy AI services, agents, credentials, or other resources that resemble legitimate, active assets in the external or internal environment, but serve no real business or operational functions. Because authorized users and workflows have no reason to engage with these resources, any interaction with them is a high-confidence indicator of unauthorized activity such as reconnaissance, or discovery. AI decoys have been shown to surface high-volume reconnaissance and actionable intelligence on threat actors' activities[[greynoise]][[zenity]].

Positioning honeypot decoys alongside genuine resources helps surface both broad and targeted recon, since an adversary mapping reachable resources in an environment will tend to engage the decoys as well as the real assets. AI decoys that are set up as enticing targets for attackers in different ways can potentially capture how an adversary probes, manipulates, or attempts to abuse them. This can reveal the adversary's intent beyond initial reconnaissance (i.e. what they attempt once engaged), as well as the tactics, techniques and procedures used by the adversary.

Interactions with decoy resources can be collected and analyzed to build a threat intelligence feed, capturing adversary source information, tooling, prompts, and techniques. This intelligence can be fed back into detection rules and defenses across the wider environment, while the placed decoys themselves can help confuse and discourage attackers (e.g. by misleading reconnaissance via presenting plausible but false targets), in the goal of wasting the adversary's efforts and resources.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0000/"><span class="relation-id">AML.T0000</span><strong>Поиск в открытых технических базах данных</strong><p>Decoy assets are catalogued by the same internet scan databases (e.g., Shodan, Censys) that adversaries query to locate exposed AI infrastructure. Adversarial engagement that originates from those platforms can grant defenders visibility into those discovery channels.</p></a>
<a class="relation-item" href="/techniques/AML.T0000.003/"><span class="relation-id">AML.T0000.003</span><strong>Базы данных с результатами сканирований</strong><p>Decoy assets are catalogued by the same internet scan databases (e.g., Shodan, Censys) that adversaries query to locate exposed AI infrastructure. Adversarial engagement that originates from those platforms can grant defenders visibility into those discovery channels.</p></a>
<a class="relation-item" href="/techniques/AML.T0006/"><span class="relation-id">AML.T0006</span><strong>Активное сканирование</strong><p>By capturing adversarial fingerprinting behavior, honeypots can help provide early warning of sweeping activity meant to discover exposed AI targets and attack surfaces being explored by adversaries.</p></a>
<a class="relation-item" href="/techniques/AML.T0006.000/"><span class="relation-id">AML.T0006.000</span><strong>Получение списка размещённых ресурсов ИИ</strong><p>By capturing adversarial fingerprinting behavior, honeypots can help provide early warning of sweeping activity meant to discover exposed AI targets and attack surfaces being explored by adversaries.</p></a>
<a class="relation-item" href="/techniques/AML.T0006.001/"><span class="relation-id">AML.T0006.001</span><strong>Query Platform Metadata APIs</strong><p>By capturing adversarial fingerprinting behavior, honeypots can help provide early warning of sweeping activity meant to discover exposed AI targets and attack surfaces being explored by adversaries.</p></a>
<a class="relation-item" href="/techniques/AML.T0006.002/"><span class="relation-id">AML.T0006.002</span><strong>Scan for Exposed AI Infrastructure</strong><p>By capturing adversarial fingerprinting behavior, honeypots can help provide early warning of sweeping activity meant to discover exposed AI targets and attack surfaces being explored by adversaries.</p></a>
<a class="relation-item" href="/techniques/AML.T0006.003/"><span class="relation-id">AML.T0006.003</span><strong>Probe AI Agent Trigger Channels</strong><p>By capturing adversarial fingerprinting behavior, honeypots can help provide early warning of sweeping activity meant to discover exposed AI targets and attack surfaces being explored by adversaries.</p></a>
<a class="relation-item" href="/techniques/AML.T0040/"><span class="relation-id">AML.T0040</span><strong>Доступ к API инференса ИИ-модели</strong><p>Honeypots capture which endpoints adversaries are targeting and how they use discovered inference APIs: the models they&#39;re interested in, the parameters they use, and the payloads they submit once they reach a live endpoint.</p></a>
<a class="relation-item" href="/techniques/AML.T0049/"><span class="relation-id">AML.T0049</span><strong>Эксплуатация приложения, доступного из интернета</strong><p>Honeypots can safely absorb real attacks and exploit attempts, turning payloads into mapped defensive signatures for a potential target without risking real assets.</p></a>
<a class="relation-item" href="/techniques/AML.T0084/"><span class="relation-id">AML.T0084</span><strong>Выявление конфигурации ИИ-агента</strong><p>Honeypots can surface which agent configuration and related assets adversaries hunt for, exposing the reconnaissance that precedes agent-targeted attacks, and acting as both canary tokens and decoy.</p></a>
<a class="relation-item" href="/techniques/AML.T0093/"><span class="relation-id">AML.T0093</span><strong>Внедрение промпта через публичное приложение</strong><p>Since honeypots capture real input, they can provide visibility into attacks and prompt payloads that are target-specific.</p></a>
<a class="relation-item" href="/techniques/AML.T0095/"><span class="relation-id">AML.T0095</span><strong>Поиск на открытых сайтах и доменах</strong><p>Decoy assets, such as honeypots, can act as a first contact for adversarial activity and reveal discovery channels that attackers are using to find web AI assets for targeting.</p></a>
<a class="relation-item" href="/techniques/AML.T0095.000/"><span class="relation-id">AML.T0095.000</span><strong>Репозитории кода</strong><p>Decoy assets, such as honeypots, can act as a first contact for adversarial activity and reveal discovery channels that attackers are using to find web AI assets for targeting.</p></a>
<a class="relation-item" href="/techniques/AML.T0132/"><span class="relation-id">AML.T0132</span><strong>Misconfigured or Publicly Exposed AI Services</strong><p>Honeypots can safely absorb real attacks and exploit attempts, turning payloads into mapped defensive signatures for a potential target without risking real assets.</p></a>
<a class="relation-item" href="/techniques/AML.T0133/"><span class="relation-id">AML.T0133</span><strong>Discover AI Agent Runtime Capabilities</strong><p>Interactive agentic decoys can present a plausible but fake toolset and capabilities. When an adversary enumerates the agent&#39;s capabilities at runtime, the honeypot can capture which tools and privileges they probe for and how they attempt to exploit it, while planted decoy credentials can also be used as canary tokens.</p></a>
</div>

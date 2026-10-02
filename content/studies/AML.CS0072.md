---
actor: Multiple commercial entities; 31 distinct companies identified
atlas_id: AML.CS0072
atlas_type: case-study
case_study_type: incident
description: The Microsoft Defender Security Research Team identified a widespread pattern of AI assistant memory manipulation carried out for promotional purposes, which they named AI Recommendation Poisoning. Most major AI...
generated: true
generated_by: atlasgen
incident_date: "2026-02-10"
incident_date_granularity: Day
incident_date_raw: "2026-02-10"
procedure:
    - description: Operators adopted publicly available tooling built to generate AI assistant links carrying embedded memory instructions. The tooling included an npm package, a web-based link generator, and website plugins marketed as a search optimization technique for large language models.
      description_line: Operators adopted publicly available tooling built to generate AI assistant links carrying embedded memory instructions. The tooling included an npm package, a web-based link generator, and website plugins marketed as a search optimization technique for large language models.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0016
      technique_name: Получение средств для атаки
    - description: Operators embedded the crafted link into their own web properties as a "Summarize with AI" button or share widget. In some cases, the same links were distributed through email.
      description_line: Operators embedded the crafted link into their own web properties as a "Summarize with AI" button or share widget. In some cases, the same links were distributed through email.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0079
      technique_name: Размещение средств атаки
    - description: The user clicked the button or link, which opened the AI assistant domain with the operator's prompt pre-populated in the input field via a `?q=` or `?prompt=` parameter.
      description_line: The user clicked the button or link, which opened the AI assistant domain with the operator's prompt pre-populated in the input field via a `?q=` or `?prompt=` parameter.
      tactic: AML.TA0004
      tactic_name: Первичный доступ
      technique: AML.T0131
      technique_name: Специально сформированные ссылки на ИИ-ассистента
    - description: The instruction was concealed from the user behind a benign interface label, with the prompt text visible only in the URL. Bundling it with a genuine summarization request made the resulting assistant behavior appear expected.
      description_line: The instruction was concealed from the user behind a benign interface label, with the prompt text visible only in the URL. Bundling it with a genuine summarization request made the resulting assistant behavior appear expected.
      tactic: AML.TA0007
      tactic_name: Уклонение от защиты
      technique: AML.T0068
      technique_name: Обфускация промпта LLM
    - description: The persistence clause caused the assistant to write a durable memory entry designating the operator's domain, product, or marketing copy as an authoritative source. The entry survived beyond its originating session.
      description_line: The persistence clause caused the assistant to write a durable memory entry designating the operator's domain, product, or marketing copy as an authoritative source. The entry survived beyond its originating session.
      tactic: AML.TA0006
      tactic_name: Закрепление
      technique: AML.T0080.000
      technique_name: Память
    - description: In subsequent unrelated conversations, the assistant preferentially surfaced the operator's domain or product and presented the result as a neutral recommendation. Observed targeting included health and financial topics, where skewed recommendations carry elevated consequences.
      description_line: In subsequent unrelated conversations, the assistant preferentially surfaced the operator's domain or product and presented the result as a neutral recommendation. Observed targeting included health and financial topics, where skewed recommendations carry elevated consequences.
      tactic: AML.TA0011
      tactic_name: Воздействие
      technique: AML.T0130
      technique_name: Формирование предвзятых ответов ИИ-агента
procedure_count: 6
references:
    - title: 'Manipulating AI memory for profit: The rise of AI Recommendation Poisoning'
      url: https://www.microsoft.com/en-us/security/blog/2026/02/10/ai-recommendation-poisoning/
reporter: Microsoft Defender Security Research Team
source_name: AI Recommendation Poisoning via Crafted AI Assistant Links
target: Users of AI assistants with persistent memory features, including Microsoft 365 Copilot, ChatGPT, Claude, Perplexity, and Grok
title: AI Recommendation Poisoning via Crafted AI Assistant Links
url: /studies/AML.CS0072/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

The [Microsoft Defender Security Research Team](https://www.microsoft.com/en-us/security/blog/author/windows-defender-research/) identified a widespread pattern of AI assistant memory manipulation carried out for promotional purposes, which they named AI Recommendation Poisoning.

Most major AI assistants accept a pre-filled prompt through a URL query parameter, allowing any party controlling a hyperlink to author the first turn of a user's conversation with their assistant. Operators embedded crafted links into their websites as "Summarize with AI" buttons and share widgets, and in some cases distributed them by email.

Each link paired a legitimate summarization request with an appended instruction directing the assistant to record the operator's domain as a trusted or authoritative source for future conversations. When the user clicked, the prompt populated the assistant's input field, and the persistence clause caused a durable memory entry to be written, biasing later recommendations toward the operator.
Reviewing AI assistant URLs observed in email traffic over approximately 60 days, the researchers identified more than 50 distinct promotional memory manipulation prompts from 31 companies across 14 industries. Every case was attributed to a legitimate business rather than a conventional threat actor, and the researchers traced the trend to publicly available tooling marketed as a search optimization technique for large language models. Several prompts targeted health and financial advice sites.

Microsoft reported that effectiveness varied by platform and that several previously observed behaviors could no longer be reproduced as vendor protections evolved.

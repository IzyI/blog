---
title: "{{ replace .Name "-" " " | title }}"  
date: {{ .Date }}  
type: it
draft: true
linkhabr: false  
description: Text about my cool site
images:
  - default_open_graph.jpg
tags: [python]  
---  
## ТЕКСТ
{{ replace .Name "-" " " | title }}  it
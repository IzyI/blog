---
title: "{{ replace .Name "-" " " | title }}"  
date: {{ .Date }}  
draft: true
type: embedder
linkhabr: false  
description: Text about my cool site
images:
  - default_open_graph.jpg
tags: [си]  

---  
## Заголовок 
{{ replace .Name "-" " " | title }}  embeder
# DIFERENÇA ENTRE CONCORRÊNCIA E PARALELISMO
Concorrência e paralelismo são dois conceitos relacionados, mas distintos, na programação concorrente e na execução de código em computadores. Eles descrevem maneiras diferentes de lidar com múltiplas tarefas ou processos simultâneos:

1. **Concorrência:**
   - Concorrência se refere à capacidade de um sistema ou programa de lidar com várias tarefas ao mesmo tempo, mesmo que não necessariamente executando-as simultaneamente. É a ideia de que várias tarefas estão sendo gerenciadas de forma eficiente, alternando entre elas para que pareça que estão acontecendo simultaneamente.
   - É mais sobre a estrutura de controle de fluxo do programa e a organização das tarefas do que sobre a execução física ao mesmo tempo.
   - É útil para lidar com problemas de concorrência, como comunicação entre threads, garantir acesso seguro a recursos compartilhados e melhorar a capacidade de resposta de aplicativos.
   - Pode ser implementado com threads, goroutines (em Go), processos ou outras abstrações de concorrência.

2. **Paralelismo:**
   - Paralelismo se refere à execução simultânea real de várias tarefas em múltiplos núcleos de CPU (ou máquinas separadas em sistemas distribuídos) para acelerar o processamento.
   - Envolve a execução real de várias partes de um programa ao mesmo tempo, geralmente em processadores separados.
   - É mais relevante quando há recursos computacionais múltiplos disponíveis, como CPU com vários núcleos.
   - O paralelismo é usado para melhorar o desempenho e a eficiência ao processar tarefas demoradas em paralelo.

Em resumo, a concorrência lida com a organização e gerenciamento de tarefas em um sistema para melhorar a eficiência e a capacidade de resposta, enquanto o paralelismo se concentra na execução simultânea real de tarefas em múltiplos núcleos ou processadores para melhorar o desempenho. Ambos são conceitos importantes na programação e computação modernas, e sua aplicação depende dos requisitos específicos do problema e dos recursos disponíveis.
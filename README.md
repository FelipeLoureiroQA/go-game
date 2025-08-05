# TEXT RPG GAME

go run .\core\.
go run .\inventario\main.go

# Fluxo

#### Jogo Primeira fase

- Criação de menu inicial (ok)
- Criação de um novo jogo 
- Você terá ervas, plantas, raizes, flores, madeiras e produtos alquimicos
- Você pode misturar e testar
- Ira criar um poção quando a combinação for correta
- Se a quantidade for incorreta podera criar reações
- Buscar receitas para as poções

# Sistemas que não seram implementado completamente mas servem de consulta

Não há uma lista definitiva de "todos os sistemas possíveis" para um RPG, pois a criatividade dos desenvolvedores é o único limite. No entanto, podemos agrupar a maioria deles em categorias lógicas para ter uma visão mais completa. A seguir, apresento uma lista detalhada dos sistemas mais comuns e de alguns mais complexos.

### Sistemas de Personagem
---
São os sistemas que definem o protagonista e sua evolução.

* **Sistema de Atributos:** Define as características básicas do personagem (Força, Destreza, Inteligência, etc.).
* **Sistema de Habilidades e Perícias:** Habilidades específicas que o personagem pode aprender e melhorar (arremesso de facas, furtividade, diplomacia, etc.).
* **Sistema de Classes:** O jogador escolhe uma "classe" que define seu estilo de jogo (Guerreiro, Mago, Ladrão).
* **Sistema de Subclasses:** Variações da classe principal, permitindo maior personalização (por exemplo, Guerreiro pode ser Cavaleiro ou Bárbaro).
* **Sistema de Progressão por Nível (XP):** O personagem ganha experiência para subir de nível, o que aumenta seus atributos e libera novas habilidades.
* **Sistema de Progressão por Habilidade:** O personagem melhora em uma habilidade específica quanto mais a usa (por exemplo, o personagem se torna um melhor espadachim ao usar mais a espada).
* **Árvore de Talentos:** Um sistema visual onde o jogador gasta pontos para desbloquear habilidades passivas ou ativas.
* **Sistema de Moralidade/Reputação:** Acompanha as decisões do jogador para determinar seu alinhamento (Bom/Mau, Ordeiro/Caótico). Isso pode afetar a interação com NPCs e o desenrolar da história.

### Sistemas de Combate
---
O núcleo da interação com inimigos.

* **Combate em Tempo Real:** Ações acontecem sem pausas, exigindo reflexos e estratégia em tempo real.
* **Combate por Turnos:** Jogador e inimigos agem em turnos. A estratégia é mais tática e deliberada.
* **Sistema de Magia:** Conjuração de feitiços usando mana, energia ou um sistema de "cooldown" (tempo de recarga).
* **Sistema de Dano e Resistências:** Calcula o dano causado por ataques, considerando tipos de dano (fogo, gelo) e resistências de inimigos.
* **Sistema de Crítico e Esquiva:** Chance de causar dano extra (crítico) ou de evitar completamente um ataque (esquiva).
* **Sistema de Combate em Equipe:** Permite gerenciar um grupo de personagens, cada um com suas próprias habilidades.

### Sistemas de Economia e Inventário
---
Gerenciamento de recursos e itens.

* **Sistema de Inventário:** A mochila do jogador onde ele armazena, organiza e gerencia itens. Pode ser por slots fixos, peso ou categorias.
* **Sistema de Equipamento:** Slots específicos para armas, armaduras, anéis e outros itens que o personagem pode vestir.
* **Sistema de Moeda:** Gerencia a moeda do jogo (ouro, moedas de prata) usada para comprar e vender itens.
* **Sistema de Criação de Itens (Crafting):** Permite combinar materiais para criar novos itens, armas ou poções.
* **Sistema de Comércio:** Permite ao jogador interagir com NPCs e outros jogadores para comprar e vender itens.

### Sistemas de Mundo
---
Como o jogador interage e explora o ambiente.

* **Sistema de Missões:** Gerencia as missões do jogo, objetivos, recompensas e o progresso do jogador.
* **Sistema de Diálogo:** Permite ao jogador conversar com NPCs, tomar decisões e afetar a narrativa.
* **Sistema de Viagem:** Mecânicas para se mover pelo mundo, como fast travel, montarias, barcos, etc.
* **Sistema de Tempo e Clima:** O jogo simula um ciclo de dia e noite, e as condições climáticas podem afetar a jogabilidade ou a estética.
* **Sistema de Facções:** Grupos no mundo do jogo com os quais o jogador pode interagir. As ações do jogador podem afetar sua reputação com essas facções.

### Sistemas Avançados
---
Sistemas mais complexos que adicionam profundidade.

* **Inteligência Artificial (IA) de NPCs:** Define como os NPCs e inimigos se comportam, reagem ao jogador e interagem entre si.
* **Sistema de Física:** Simula a física do mundo para tornar as interações mais realistas e dinâmicas (ex: objetos que caem, projéteis que seguem uma trajetória).
* **Sistema de Geração Procedural:** O conteúdo do jogo (mapas, masmorras, itens) é gerado de forma aleatória por um algoritmo, tornando cada jogatina única.
* **Sistema de Economia Dinâmica:** O preço dos itens muda com base na oferta e na demanda, tornando o mercado mais realista.
* **Sistema de Afeto/Relacionamento:** As interações com NPCs podem levar a amizades, rivalidades ou relacionamentos amorosos, que podem afetar o enredo.

Esses sistemas podem ser combinados de inúmeras maneiras para criar jogos com experiências completamente diferentes. Por exemplo, um RPG como *Elden Ring* foca em um combate de ação em tempo real e exploração de mundo, enquanto um jogo como *Divinity: Original Sin 2* se concentra mais em combate por turnos e em um sistema de diálogos robusto.
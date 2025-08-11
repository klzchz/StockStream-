# 📈 BrokerFlow

**BrokerFlow** é um sistema de **Home Broker** em tempo real, projetado para simular negociações no mercado financeiro com alta performance.  
A arquitetura integra múltiplas tecnologias modernas, garantindo escalabilidade, baixa latência e comunicação eficiente entre serviços.

---

## 🚀 Tecnologias Utilizadas

- **Frontend**: React + Next.js → Interface dinâmica e responsiva para o usuário.
- **Backend**: Nest.js → API REST robusta e modular.
- **Mensageria**: Apache Kafka → Processamento assíncrono e escalável de eventos.
- **Motor de Bolsa**: Golang → Execução rápida de ordens simuladas e gerenciamento de cotações.
- **Comunicação**: REST (sincrono) + Kafka (assíncrono).

---

## 📊 Arquitetura

```mermaid
graph LR
    A[User] <--> B[Home Broker (React/Next.js)]
    B <--> |REST| C[Nest.js (Backend)]
    C --> D[Apache Kafka]
    D <--> E[Sistema Bolsa (Golang)]
Fluxo de funcionamento:

O usuário interage com o Home Broker via interface web.

O frontend envia requisições REST para o backend Nest.js.

O backend publica e consome eventos no Apache Kafka.

O Sistema Bolsa (em Golang) processa ordens e envia atualizações de volta via Kafka.

As atualizações retornam ao usuário em tempo real.

🛠 Como Executar
1️⃣ Pré-requisitos
Node.js >= 18

Golang >= 1.21

Docker e Docker Compose

2️⃣ Clonar repositório

git clone https://github.com/seuusuario/brokerflow.git
cd brokerflow
3️⃣ Subir infraestrutura
bash
Copiar
Editar
docker-compose up -d
4️⃣ Executar serviços
Backend Nest.js

bash
Copiar
Editar
cd backend
npm install
npm run start:dev
Frontend React/Next.js

bash
Copiar
Editar
cd frontend
npm install
npm run dev
Sistema Bolsa Golang

bash
Copiar
Editar
cd sistema-bolsa
go run main.go
💡 Diferenciais
Arquitetura distribuída e escalável.

Uso de mensageria para desacoplamento entre serviços.

Simulação realista de ordens de compra e venda.

Multi-stack (JavaScript/TypeScript + Golang).


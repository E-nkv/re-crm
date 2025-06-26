# re-crm

Welcome to **re-crm**, a modern, robust Customer Relationship Management system designed to empower your team with seamless collaboration, fine-grained access control, and an intuitive user experience.

And yes, the README was mostly done by GPT since real 10x devs don't write READMEs 🤑
---

## 🌟 Features

- **Role-Based Access Control (RBAC)**  
  Tailored access for different user roles including **Superadmins**, **Admins**, **Managers**, and **Sales Representatives**. Each role has customized permissions ensuring security and appropriate data visibility.

- **Personalized Dashboards**  
  Every user is greeted with a dynamic, personalized dashboard.

- **Real-Time Communication**  
  Stay connected effortlessly with integrated chat functionality:  
  - **Direct Messaging:** Active users can chat one-on-one for quick, private conversations.  
  - **Channel-Based Discussions:** Create and participate in channels, similar to WhatsApp groups, to collaborate on projects, share updates, and foster team synergy. Supports img uploading. 

---

## 🛠️ Tech Stack

- **Backend API:**  
  Built with [Go](https://golang.org/) leveraging the [Chi](https://github.com/go-chi/chi) router for lightweight, modular HTTP services and [GORM](https://gorm.io/) as the ORM for seamless database interactions.

- **Database:**  
  Powered by [PostgreSQL](https://www.postgresql.org/), a powerful, reliable, and scalable relational database system that ensures your data is safe and performant.

- **Frontend:**  
  Developed using [React](https://reactjs.org/) with **TypeScript** for type safety and maintainability, styled elegantly with [Tailwind CSS](https://tailwindcss.com/) for rapid, responsive UI development.

- **Containerization & Deployment:**  
  Utilizes [Docker](https://www.docker.com/) to streamline development and deployment workflows:  
  - In development, Docker spins up the PostgreSQL database 
  - In production, Docker orchestrates the deployment of the entire application stack, ensuring consistency and scalability.

---

## Want to collaborate?

Open a PR with your suggested changes and I'll be happy to look into it


<!-- ## 🚀 Getting Started

To get a local development environment up and running:

 -->
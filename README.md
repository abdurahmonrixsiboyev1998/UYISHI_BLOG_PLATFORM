# Uy vazifasi: Oddiy Blog Platformasi

## Topshiriq Tavsifi:
Foydalanuvchilar ro'yxatdan o'tishi, autentifikatsiyadan o'tishi va blog postlarini boshqarishi (`CRUD`) mumkin bo'lgan oddiy blog platformasini yaratish.


## Talablar:
1. **Foydalanuvchi autentifikatsiyasi:**
    - Foydalanuvchini ro'yxatdan o'tkazish (`POST /register`).
    - Kirish funksiyasini yaratish (`POST /login`), bu `JWT` tokenini qaytarishi kerak.

2. **Blog postlarini boshqarish (CRUD)::**
    - Blog post yaratish: Avtorizatsiya qilingan foydalanuvchilar blog postlarini yaratishlari mumkin (`POST /posts`).
    - Barcha postlarni olish: Foydalanuvchilar barcha postlarni ko'rishlari mumkin (`GET /posts`).
    - Blog postni yangilash: Foydalanuvchilar o'z blog postlarini yangilashlari mumkin (`PUT /posts/{id}`).
    - Blog postni o'chirish: Foydalanuvchilar o'z blog postlarini o'chirishlari mumkin (`DELETE /posts/{id}`).
    
3. **Markdown qo'llab-quvvatlovi:**
    - Blog kontenti `Markdown` formatida bo'lishi va uni to'g'ri ko'rsatish kerak.

4. **Ma'lumotlar bazasi sozlamalari:**
    - `PostgreSQL`da kerakli jadvallarni yarating:
        - `users` jadvali foydalanuvchi ma'lumotlarini saqlash uchun (`id`, `foydalanuvchi nomi`, `email`, `parol xesh`).
        - `posts` jadvali blog postlarini saqlash uchun (`id`, `user_id`, `title`, `content`, `created_at`, `updated_at`).

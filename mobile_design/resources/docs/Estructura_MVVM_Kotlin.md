## 📂 Ejemplo de estructura **simple y escalable** (MVVM reducido)

```markdown
MyApp/
├── app/
│   ├── src/
│   │   ├── main/java/com/example/myapp/
│   │   │   ├── data/                # Repositorios y acceso a datos
│   │   │   │   ├── UserRepository.kt
│   │   │   │   └── UserRepository_test.kt
│   │   │   ├── domain/              # Casos de uso / lógica de negocio
│   │   │   │   ├── GetUserUseCase.kt
│   │   │   │   └── GetUserUseCase_test.kt
│   │   │   ├── ui/                  # Views + ViewModels juntos
│   │   │   │   ├── MainActivity.kt
│   │   │   │   ├── MainViewModel.kt
│   │   │   │   └── MainViewModel_test.kt
│   │   │   ├── utils/               # Helpers y extensiones
│   │   │   │   ├── Extensions.kt
│   │   │   │   └── Extensions_test.kt
│   │   │   └── di/                  # Inyección de dependencias (si se usa)
│   │   │       └── AppModule.kt
│   │   ├── androidTest/             # Tests instrumentados (UI)
│   │   └── res/                     # Layouts, strings, drawables
│   └── build.gradle
├── docs/                            # Documentación y guías
│   └── architecture.md
├── benchmarks/                      # Benchmarks (si aplica)
└── README.md
```

---

### 🔹 Claves de esta estructura

1. **Menos carpetas** → agrupamos por capa principal (`data`, `domain`, `ui`) y evitamos subdividir demasiado.
2. **Tests pegados al código** → igual que en Go, basta con poner `_test.kt` junto al archivo principal.

   * Ejemplo: `UserRepository.kt` y `UserRepository_test.kt`.
3. **Escalable** → si el proyecto crece, puedes dividir en subcarpetas dentro de `ui/` (ej. `login/`, `profile/`).
4. **Documentación externa** en `docs/` para que los desarrolladores tengan reglas claras.

---

### 🔹 Ventaja frente a la estructura gigante

* Más **fácil de navegar** para un equipo pequeño o mediano.
* Evita “carpetitis” (tener carpetas con un solo archivo).
* Mantiene la **separación de responsabilidades** sin complicar.

---

👉 En Android, muchos equipos grandes usan **MVVM o Clean Architecture con muchas carpetas**, pero si vienes del mundo Go, esta aproximación **minimalista + escalable** encaja mejor con tu mentalidad.

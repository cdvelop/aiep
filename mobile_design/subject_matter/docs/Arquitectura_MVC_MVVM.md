## 🏗️ Modelos de Arquitectura: MVC vs. MVVM

Son patrones para organizar tu código, separar responsabilidades y hacerlo más mantenible.

### Modelo-Vista-Controlador (MVC)

Es el patrón clásico, muy común en desarrollo web.

  * **Modelo**: Los datos y la lógica de negocio (ej: una struct `Usuario` en Go).
  * **Vista**: La interfaz de usuario (ej: una plantilla HTML).
  * **Controlador**: Recibe las peticiones del usuario, interactúa con el Modelo y decide qué Vista mostrar.

**Flujo**: `Usuario -> Controlador -> Modelo -> Vista -> Usuario`

```go
// Ejemplo conceptual en Go (para una API REST)

// --- MODELO ---
type Usuario struct {
    Nombre string
    Email  string
}

// --- CONTROLADOR ---
func obtenerUsuario(w http.ResponseWriter, r *http.Request) {
    // 1. El controlador recibe la petición.
    id := r.URL.Query().Get("id")

    // 2. Interactúa con el modelo (aquí simulamos obtener datos).
    usuario := Usuario{Nombre: "Ada Lovelace", Email: "ada@example.com"}

    // 3. Prepara la "Vista" (en este caso, datos JSON).
    json.NewEncoder(w).Encode(usuario)
}
```

### Modelo-Vista-ViewModel (MVVM)

Muy popular en desarrollo móvil moderno (especialmente con Kotlin y Swift).

  * **Modelo**: Igual que en MVC, los datos.
  * **Vista**: La UI. Es "tonta", solo se encarga de mostrar datos y notificar acciones del usuario.
  * **ViewModel**: Es el intermediario. Expone los datos del Modelo de una forma que la Vista pueda usar fácilmente y contiene la lógica de presentación. **No tiene referencia directa a la Vista**.

**Flujo**: `Usuario <-> Vista <-> ViewModel <-> Modelo`

La clave es el **"Data Binding"**: la Vista se "suscribe" a los cambios en el ViewModel y se actualiza automáticamente.

```kotlin
// Ejemplo conceptual en Kotlin (para Android)

// --- MODELO ---
data class Usuario(val nombre: String, val email: String)

// --- VIEWMODEL ---
class UsuarioViewModel : ViewModel() {
    // LiveData es un "contenedor de datos observable".
    val nombreUsuario: MutableLiveData<String> by lazy {
        MutableLiveData<String>()
    }
    
    fun cargarUsuario() {
        // Lógica para obtener datos del Modelo...
        val usuario = Usuario("Grace Hopper", "grace@example.com")
        nombreUsuario.value = usuario.nombre // Actualiza el LiveData
    }
}

// --- VISTA (Activity/Fragment) ---
class PerfilActivity : AppCompatActivity() {
    private lateinit var viewModel: UsuarioViewModel

    override fun onCreate(savedInstanceState: Bundle?) {
        // ...
        viewModel = ViewModelProvider(this).get(UsuarioViewModel::class.java)

        // La Vista "observa" los cambios en el ViewModel.
        viewModel.nombreUsuario.observe(this, Observer { nombre ->
            // Cuando el nombre cambie en el ViewModel, este código se ejecuta.
            textViewNombre.text = nombre
        })

        viewModel.cargarUsuario()
    }
}
```

**Diferencia clave**: En **MVVM**, el ViewModel no conoce a la Vista. Simplemente emite datos. La Vista decide cómo reaccionar a esos datos. Esto desacopla mucho el código y facilita las pruebas.

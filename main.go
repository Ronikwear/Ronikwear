package main

import "net/http"

// Обработчик для создания новой заметки.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		// Используем функцию http.Error() для отправки кода состояния 405 с соответствующим сообщением.
		http.Error(w, "Метод запрещен!", 405)
		return
	}

	w.Write([]byte("Создание новой заметки..."))
}

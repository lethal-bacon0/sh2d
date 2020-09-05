build:
	go build -o bin/asteroids main.go

android:
	mkdir -p ./mobile/android/asteroids/
	go run github.com/hajimehoshi/ebiten/cmd/ebitenmobile bind -target android -javapkg com.github.lethalbacon0.asteroids -o ./mobile/android/app/libs/asteroids.aar ./mobile

clean:
	rm -rf ./bin
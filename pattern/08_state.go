/*
Состояние - это поведенческий паттерн проектирования, который позволяет объектам менять поведение в зависимости от своего состояния.
 Извне создаётся впечатление, что изменился класс объекта.

Применимость:
- Когда у вас есть объект, поведение которого кардинально меняется в зависимости от внутреннего состояния, причём типов состояний много, и их код часто меняется.
- Когда код класса содержит множество больших, похожих друг на друга, условных операторов, которые выбирают поведения в зависимости от текущих значений полей класса.
- Когда вы сознательно используете табличную машину состояний, построенную на условных операторах, но
вынуждены мириться с дублированием кода для похожих состояний и переходов.
Плюсы:
- Избавляет от множества больших условных операторов машины состояний.
- Концентрирует в одном месте код, связанный с определённым состоянием.
- Упрощает код контекста.
Минусы:
- Может неоправданно усложнить код, если состояний мало и они редко меняются.
*/

package main

import (
	"fmt"
)

type AudioPlayer struct {
	UnlockedScreen State
	LockedScreen   State
	state          State
	volume         int
	track          string
}

type State interface {
	audionext() string
	audioplay() string
	audioprevious() string
}

func newAudioPlayer(volume int, track string) *AudioPlayer {
	a := &AudioPlayer{
		volume: volume,
		track:  track,
	}
	LockedScreenState := &LockedScreen{
		Audioplayer: a,
	}
	UnlockedScreenState := &UnlockedScreen{
		AudioPlayer: a,
	}

	a.setState(LockedScreenState)
	// a.setState()
	a.UnlockedScreen = UnlockedScreenState
	a.LockedScreen = LockedScreenState

	return a
}
func (a *AudioPlayer) audionext() string {
	return a.state.audionext()
}

func (a *AudioPlayer) audioplay() string {
	return a.state.audioplay()
}

func (a *AudioPlayer) audioprevious() string {
	return a.state.audioprevious()
}

func (a *AudioPlayer) setState(s State) {
	a.state = s
}

type LockedScreen struct {
	Audioplayer *AudioPlayer
}

func (n *LockedScreen) audionext() string {
	state := "Screen is locked, can't set next track"
	fmt.Println(state)
	return state
}

func (n *LockedScreen) audioplay() string {
	fmt.Println(n.Audioplayer.track)
	return n.Audioplayer.track
}

func (n *LockedScreen) audioprevious() string {
	fmt.Println("screen is locked, can't set previous track")
	return n.Audioplayer.track
}

type UnlockedScreen struct {
	AudioPlayer *AudioPlayer
}

func (u *UnlockedScreen) audionext() string {
	fmt.Println("Set new track ", u.AudioPlayer.track)
	return u.AudioPlayer.track

}

func (u *UnlockedScreen) audioplay() string {
	fmt.Println("track is ", u.AudioPlayer.track)
	return u.AudioPlayer.track

}

func (u *UnlockedScreen) audioprevious() string {
	fmt.Println("Set previous track is")
	return u.AudioPlayer.track
}

func main() {
	audioPlayer := newAudioPlayer(100, "Track 2")

	v2 := audioPlayer.audioprevious()
	fmt.Println(v2)

	v1 := audioPlayer.audioplay()
	fmt.Println(v1)

}

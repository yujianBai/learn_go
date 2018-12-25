package library

import (
	"testing"
)

func TestOps(t *testing.T){
	mm := NewMusicManager()
	if mm == nil{
		t.Error("NewMusicManage failed.")
	}
	if mm.len() != 0{
		t.Error("NewMusicMange failed, not empty")
	}

	m0 := &MusicEntry{"1", "My Heart Will Go On", "Celion Dion", "Pop", "http://qbox.me/24501234", "mp3"}
	mm.Add(m0)
	if mm.Len() == 0{
		t.Error("MusicManger.Add() filed.")
	}

	m,_ := mm.Find(m0.Name)
	if m == nil{
		t.Error("MusicManger.Find() filed.")
	}

	if m.Id != m0.Id || m.Artist != m0.Artist || m.Name != m0.Name || m.Genre != m0.Genre ||
	m.Source != m0.Source || m.Type != m0.Type{
		t.Error("MusicManage.Find() filed. Found item mismatch")
	}

	m, err := mm.Get(0)
	if m == nil{
		t.Error("MusicManage.Get() filed.", err)
	}

	m = mm.Remove(0)
	if m == nil{
		t.Error("MusicMange.Remove() failed.", err)
	}

}
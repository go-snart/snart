package route

import "testing"

func TestWait(t *testing.T) {
	_, _, _, _, _, _,
		c := ctxDummy("owo")
	w := c.Wait(True, True)

	if w == nil {
		t.Fatal("w == nil")
	}
	if val(w.general) != val(True) {
		t.Fatal("w.general != True")
	}
	if val(w.specific) != val(True) {
		t.Fatal("w.specific != True")
	}
	if w.Return == nil {
		t.Fatal("w.Return == nil")
	}
	if w.cancel == nil {
		t.Fatal("w.cancel == nil")
	}
}

func TestWaitNoCancel(t *testing.T) {
	_, _, _, _, _, _,
		c := ctxDummy("owo")
	w := c.WaitCancel(True, True, false)

	if w.cancel != nil {
		t.Fatal("w.cancel != nil")
	}
}

func TestWaitHandle(t *testing.T) {
	_, _, _, _, _, _,
		c := ctxDummy("owo")
	w := c.Wait(True, True)

	_,
		mc := messageCreateDummy("uwu")
	go w.handle(c.Session, mc)

	nc := <-w.Return
	if mc.Message != nc.Message {
		t.Fatal("mc.Message != nc.Message")
	}
}

func TestWaitHandleNoGeneral(t *testing.T) {
	_, _, _, _, _, _,
		c := ctxDummy("owo")
	w := c.Wait(False, True)

	_,
		mc := messageCreateDummy("uwu")
	w.handle(c.Session, mc)
}

func TestWaitHandleNoSpecific(t *testing.T) {
	_, _, _, _, _, _,
		c := ctxDummy("owo")
	w := c.Wait(True, False)

	_,
		mc := messageCreateDummy("uwu")
	go w.handle(c.Session, mc)

	nc := <-w.Return
	if nc != nil {
		t.Fatal("nc != nil")
	}
}

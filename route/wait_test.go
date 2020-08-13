package route_test

/*
func TestWait(t *testing.T) {
	_, _, _, _, _,
		c := ctxDummy("owo")
	w := c.Wait(route.True, route.True)

	switch {
	case w == nil:
		t.Fatal("w == nil")
	case val(w.General) != val(route.True):
		t.Fatal("w.General != route.True")
	case val(w.Specific) != val(route.True):
		t.Fatal("w.Specific != route.True")
	case w.Return == nil:
		t.Fatal("w.Return == nil")
	case w.Cancel == nil:
		t.Fatal("w.Cancel == nil")
	}
}

func TestWaitNoCancel(t *testing.T) {
	_, _, _, _, _,
		c := ctxDummy("owo")
	w := c.WaitCancel(route.True, route.True, false)

	if w.Cancel != nil {
		t.Fatal("w.Cancel != nil")
	}
}

func TestWaitHandle(t *testing.T) {
	_, _, _, _, _,
		c := ctxDummy("owo")
	w := c.Wait(route.True, route.True)

	_,
		mc := messageCreateDummy("uwu")
	go w.Handle(c.Session, mc)

	nc := <-w.Return
	if mc.Message != nc.Message {
		t.Fatal("mc.Message != nc.Message")
	}
}

func TestWaitHandleNoGeneral(t *testing.T) {
	_, _, _, _, _,
		c := ctxDummy("owo")
	w := c.Wait(route.False, route.True)

	_,
		mc := messageCreateDummy("uwu")
	w.Handle(c.Session, mc)
}

func TestWaitHandleNoSpecific(t *testing.T) {
	_, _, _, _, _,
		c := ctxDummy("owo")
	w := c.Wait(route.True, route.False)

	_,
		mc := messageCreateDummy("uwu")
	go w.Handle(c.Session, mc)

	nc := <-w.Return
	if nc != nil {
		t.Fatal("nc != nil")
	}
}
*/

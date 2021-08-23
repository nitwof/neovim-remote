package cmder

type Cmder struct {
	Exec func(args ...string) error
}

func (c Cmder) InstallAppimage(name, url, dst string) error {

	return nil
}

func (c Cmder) ExtractAppimage(file, dst string) error {
	var err error

	err = c.Exec("chmod", "u+x", file)
	if err != nil {
		return err
	}

	err = c.Exec(file, "--appimage-extract")
	if err != nil {
		return err
	}

	return nil
}

func (c Cmder) Download(url, dst string) error {
	return c.Exec("curl", "-L", url, "-o", dst)
}

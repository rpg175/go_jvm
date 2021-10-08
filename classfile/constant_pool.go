package classfile

type ConstantPool []ConstantInfo

func readContentPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo,cpCount)

	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader,cp)
		switch cp[i].(type) {
		case *ConstantsLongInfo, *ConstantDoubleInfo:
			i++
		}
	}

	return cp
}



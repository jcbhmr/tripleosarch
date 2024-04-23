package tripleosarch

type Triple struct {}

type TripleArchType = string

const (
	TripleArchTypeUnknownArch = iota,
    TripleArchTypeArm,            
    TripleArchTypeArmeb,          
    TripleArchTypeAarch64,        
    TripleArchTypeAarch64_be,     
    TripleArchTypeAarch64_32,     
    TripleArchTypeArc,            
    TripleArchTypeAvr,            
    TripleArchTypeBpfel,          
    TripleArchTypeBpfeb,          
    TripleArchTypeCsky,           
    TripleArchTypeDxil,           
    TripleArchTypeHexagon,        
    TripleArchTypeLoongarch32,    
    TripleArchTypeLoongarch64,    
    TripleArchTypeM68k,           
    TripleArchTypeMips,           
    TripleArchTypeMipsel,         
    TripleArchTypeMips64,         
    TripleArchTypeMips64el,       
    TripleArchTypeMsp430,         
    TripleArchTypePpc,            
    TripleArchTypePpcle,          
    TripleArchTypePpc64,          //
    TripleArchTypePpc64le,        
    TripleArchTypeR600,           
    TripleArchTypeAmdgcn,         
    TripleArchTypeRiscv32,        
    TripleArchTypeRiscv64,        
    TripleArchTypeSparc,          
    TripleArchTypeSparcv9,        
    TripleArchTypeSparcel,        
    TripleArchTypeSystemz,        
    TripleArchTypeTce,            
    TripleArchTypeTcele,          
    TripleArchTypeThumb,          
    TripleArchTypeThumbeb,        
    TripleArchTypeX86,            
    TripleArchTypeX86_64,         
    TripleArchTypeXcore,          
    TripleArchTypeXtensa,         
    TripleArchTypeNvptx,          
    TripleArchTypeNvptx64,        
    TripleArchTypeLe32,           
    TripleArchTypeLe64,           
    TripleArchTypeAmdil,          
    TripleArchTypeAmdil64,        
    TripleArchTypeHsail,          
    TripleArchTypeHsail64,        
    TripleArchTypeSpir,           
    TripleArchTypeSpir64,         
    TripleArchTypeSpirv,          
    TripleArchTypeSpirv32,        
    TripleArchTypeSpirv64,        
    TripleArchTypeKalimba,        
    TripleArchTypeShave,          
    TripleArchTypeLanai,          
    TripleArchTypeWasm32,         
    TripleArchTypeWasm64,         
    TripleArchTypeRenderscript32, 
    TripleArchTypeRenderscript64, 
    TripleArchTypeVe,             
    TripleArchTypeLastArchType = TripleArchTypeVe
)

func (t TripleArchType) String() string {
	switch t {
		case UnknownArch:    return "unknown";
	  
		case TripleArchTypeAarch64:        return "aarch64";
		case TripleArchTypeAarch64_32:     return "aarch64_32";
		case TripleArchTypeAarch64_be:     return "aarch64_be";
		case amdgcn:         return "amdgcn";
		case amdil64:        return "amdil64";
		case amdil:          return "amdil";
		case TripleArchTypeArc:            return "arc";
		case TripleArchTypeArm:            return "arm";
		case TripleArchTypeArmeb:          return "armeb";
		case TripleArchTypeAvr:            return "avr";
		case TripleArchTypeBpfeb:          return "bpfeb";
		case TripleArchTypeBpfel:          return "bpfel";
		case TripleArchTypeCsky:           return "csky";
		case TripleArchTypeDxil:           return "dxil";
		case TripleArchTypeHexagon:        return "hexagon";
		case hsail64:        return "hsail64";
		case hsail:          return "hsail";
		case kalimba:        return "kalimba";
		case lanai:          return "lanai";
		case le32:           return "le32";
		case le64:           return "le64";
		case TripleArchTypeLoongarch32:    return "loongarch32";
		case TripleArchTypeLoongarch64:    return "loongarch64";
		case TripleArchTypeM68k:           return "m68k";
		case TripleArchTypeMips64:         return "mips64";
		case TripleArchTypeMips64el:       return "mips64el";
		case TripleArchTypeMips:           return "mips";
		case TripleArchTypeMipsel:         return "mipsel";
		case TripleArchTypeMsp430:         return "msp430";
		case nvptx64:        return "nvptx64";
		case nvptx:          return "nvptx";
		case TripleArchTypePpc64:          return "powerpc64";
		case ppc64le:        return "powerpc64le";
		case TripleArchTypePpc:            return "powerpc";
		case TripleArchTypePpcle:          return "powerpcle";
		case r600:           return "r600";
		case renderscript32: return "renderscript32";
		case renderscript64: return "renderscript64";
		case riscv32:        return "riscv32";
		case riscv64:        return "riscv64";
		case shave:          return "shave";
		case sparc:          return "sparc";
		case sparcel:        return "sparcel";
		case sparcv9:        return "sparcv9";
		case spir64:         return "spir64";
		case spir:           return "spir";
		case spirv:          return "spirv";
		case spirv32:        return "spirv32";
		case spirv64:        return "spirv64";
		case systemz:        return "s390x";
		case tce:            return "tce";
		case tcele:          return "tcele";
		case thumb:          return "thumb";
		case thumbeb:        return "thumbeb";
		case ve:             return "ve";
		case wasm32:         return "wasm32";
		case wasm64:         return "wasm64";
		case x86:            return "i386";
		case x86_64:         return "x86_64";
		case xcore:          return "xcore";
		case xtensa:         return "xtensa";
	}
	panic("Invalid ArchType!")
}
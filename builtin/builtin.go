// Built-in functions
package builtin

import (
	"fmt"
	"github.com/ncw/gpython/py"
)

const builtin_doc = `Built-in functions, exceptions, and other objects.

Noteworthy: None is the 'nil' object; Ellipsis represents '...' in slices.`

// Initialise the module
func init() {
	methods := []*py.Method{
		// py.NewMethod("__build_class__", builtin___build_class__, 0, build_class_doc),
		// py.NewMethod("__import__", builtin___import__, 0, import_doc),
		py.NewMethod("abs", builtin_abs, 0, abs_doc),
		// py.NewMethod("all", builtin_all, 0, all_doc),
		// py.NewMethod("any", builtin_any, 0, any_doc),
		// py.NewMethod("ascii", builtin_ascii, 0, ascii_doc),
		// py.NewMethod("bin", builtin_bin, 0, bin_doc),
		// py.NewMethod("callable", builtin_callable, 0, callable_doc),
		// py.NewMethod("chr", builtin_chr, 0, chr_doc),
		// py.NewMethod("compile", builtin_compile, 0, compile_doc),
		// py.NewMethod("delattr", builtin_delattr, 0, delattr_doc),
		// py.NewMethod("dir", builtin_dir, 0, dir_doc),
		// py.NewMethod("divmod", builtin_divmod, 0, divmod_doc),
		// py.NewMethod("eval", builtin_eval, 0, eval_doc),
		// py.NewMethod("exec", builtin_exec, 0, exec_doc),
		// py.NewMethod("format", builtin_format, 0, format_doc),
		// py.NewMethod("getattr", builtin_getattr, 0, getattr_doc),
		// py.NewMethod("globals", builtin_globals, py.METH_NOARGS, globals_doc),
		// py.NewMethod("hasattr", builtin_hasattr, 0, hasattr_doc),
		// py.NewMethod("hash", builtin_hash, 0, hash_doc),
		// py.NewMethod("hex", builtin_hex, 0, hex_doc),
		// py.NewMethod("id", builtin_id, 0, id_doc),
		// py.NewMethod("input", builtin_input, 0, input_doc),
		// py.NewMethod("isinstance", builtin_isinstance, 0, isinstance_doc),
		// py.NewMethod("issubclass", builtin_issubclass, 0, issubclass_doc),
		// py.NewMethod("iter", builtin_iter, 0, iter_doc),
		// py.NewMethod("len", builtin_len, 0, len_doc),
		// py.NewMethod("locals", builtin_locals, py.METH_NOARGS, locals_doc),
		// py.NewMethod("max", builtin_max, 0, max_doc),
		// py.NewMethod("min", builtin_min, 0, min_doc),
		// py.NewMethod("next", builtin_next, 0, next_doc),
		// py.NewMethod("oct", builtin_oct, 0, oct_doc),
		// py.NewMethod("ord", builtin_ord, 0, ord_doc),
		py.NewMethod("pow", builtin_pow, 0, pow_doc),
		py.NewMethod("print", builtin_print, 0, print_doc),
		// py.NewMethod("repr", builtin_repr, 0, repr_doc),
		py.NewMethod("round", builtin_round, 0, round_doc),
		// py.NewMethod("setattr", builtin_setattr, 0, setattr_doc),
		// py.NewMethod("sorted", builtin_sorted, 0, sorted_doc),
		// py.NewMethod("sum", builtin_sum, 0, sum_doc),
		// py.NewMethod("vars", builtin_vars, 0, vars_doc),
	}
	py.NewModule("builtins", builtin_doc, methods)
}

const print_doc = `print(value, ..., sep=' ', end='\\n', file=sys.stdout, flush=False)

Prints the values to a stream, or to sys.stdout by default.
Optional keyword arguments:
file:  a file-like object (stream); defaults to the current sys.stdout.
sep:   string inserted between values, default a space.
end:   string appended after the last value, default a newline.
flush: whether to forcibly flush the stream.`

func builtin_print(self py.Object, args py.Tuple, kwargs py.StringDict) py.Object {
	fmt.Printf("print %v, %v, %v\n", self, args, kwargs)
	return py.None
}

const pow_doc = `pow(x, y[, z]) -> number

With two arguments, equivalent to x**y.  With three arguments,
equivalent to (x**y) % z, but may be more efficient (e.g. for ints).`

func builtin_pow(self py.Object, args py.Tuple) py.Object {
	var v, w, z py.Object
	z = py.None
	py.UnpackTuple(args, "pow", 2, 3, &v, &w, &z)
	return py.Pow(v, w, z)
}

const abs_doc = `"abs(number) -> number

Return the absolute value of the argument.`

func builtin_abs(self, v py.Object) py.Object {
	return py.Abs(v)
}

const round_doc = `round(number[, ndigits]) -> number

Round a number to a given precision in decimal digits (default 0 digits).
This returns an int when called with one argument, otherwise the
same type as the number. ndigits may be negative.`

func builtin_round(self py.Object, args py.Tuple, kwargs py.StringDict) py.Object {
	var number, ndigits py.Object
	ndigits = py.Int(0)
	// var kwlist = []string{"number", "ndigits"}
	// FIXME py.ParseTupleAndKeywords(args, kwargs, "O|O:round", kwlist, &number, &ndigits)
	py.UnpackTuple(args, "round", 1, 2, &number, &ndigits)

	numberRounder, ok := number.(py.I__round__)
	if !ok {
		// FIXME TypeError
		panic(fmt.Sprintf("TypeError: type %s doesn't define __round__ method", number.Type().Name))
	}

	return numberRounder.M__round__(ndigits)
}